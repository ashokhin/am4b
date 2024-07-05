package com.ashokhin.am4bot.bot;

import java.time.Duration;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.ArrayList;
import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.openqa.selenium.By;
import org.openqa.selenium.ElementNotInteractableException;
import org.openqa.selenium.PageLoadStrategy;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;
import org.openqa.selenium.support.ui.ExpectedConditions;
import org.openqa.selenium.support.ui.Select;
import org.openqa.selenium.support.ui.WebDriverWait;

import com.ashokhin.am4bot.model.APIXpath;
import com.google.common.base.CharMatcher;

public class BotBase implements Runnable {
    private static final Logger logger = LogManager.getLogger(BotBase.class);
    private static final long LOGIN_TIMEOUT_SEC = 30;
    private static final int MAX_LOGIN_ATTEMPTS = 5;
    private static final long MAX_TIME_DELTA_SEC = 60;
    private static final List<String> GOOGLE_PARAMETERS = new ArrayList<String>() {
        {
            add("--headless");
            add("--no-sandbox");
            add("--disable-dev-shm-usage");
        }
    };

    private int loginAttempts = 0;
    private Instant loginLastAttemptTimestamp = Instant.now();
    private boolean isLoggedIn = false;
    private String login;
    private String password;
    private String baseURL;
    private WebDriver webDriver;

    public BotBase(String baseUrl, String login, String password) {
        this.baseURL = baseUrl;
        this.login = login;
        this.password = password;
    }

    @Override
    public void run() {
    }

    /**
     * Set options for Google Chrome
     * 
     * @param arguments The arguments to use when starting Chrome.
     */
    private final ChromeOptions setChromeOptions(List<String> arguments) {
        logger.info("Set chrome options");

        ChromeOptions chromeOptions = new ChromeOptions();

        for (String argument : arguments) {
            logger.debug(String.format("Set Google Chrome argument '%s'", argument));

            chromeOptions.addArguments(argument);
        }

        chromeOptions.setPageLoadStrategy(PageLoadStrategy.NORMAL);

        return chromeOptions;
    }

    /** Click button which found by xPath */
    protected void clickButton(String buttonXpath) {
        logger.trace(String.format("Click button '%s' as xPath", buttonXpath));
        try {
            this.webDriver.findElement(By.xpath(buttonXpath)).click();
        } catch (ElementNotInteractableException e) {
            logger.warn(String.format("Button '%s' not active", buttonXpath));
        }

        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }

    /** Click button which found as WebElement */
    protected void clickButton(WebElement webElement) {
        logger.trace(String.format("Click button '%s' as WebElement", webElement.toString()));
        webElement.click();
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }

    /** Type given text in text field found by xPath */
    protected void typeTextInField(String textFieldXpath, String enteredText) {
        logger.trace(String.format("Enter '%s' in text field '%s'", enteredText, textFieldXpath));
        this.webDriver.findElement(By.xpath(textFieldXpath)).clear();
        this.webDriver.findElement(By.xpath(textFieldXpath)).sendKeys(enteredText);
    }

    /** Find and return text from text field which given as xPath */
    protected String getTextFromElement(String elementXpath) {
        logger.trace(String.format("Get text from element '%s'", elementXpath));

        return this.webDriver.findElement(By.xpath(elementXpath)).getText();
    }

    /** Find and return text from text field which given as xPath */
    protected String getTextFromElement(WebElement webElement) {
        logger.trace(String.format("Get text from element '%s'", webElement.toString()));

        return webElement.getText();
    }

    protected int getIntFromElement(String elementXpath) {
        logger.trace(String.format("Get int from element '%s'", elementXpath));
        String elementText = this.getTextFromElement(elementXpath);
        logger.trace(String.format("Got text '%s' from element '%s'", elementText, elementXpath));

        return Integer.parseInt(
                CharMatcher.inRange('0', '9').retainFrom(elementText));
    }

    protected int getIntFromElement(WebElement webElement) {
        logger.trace(String.format("Get int from element '%s'", webElement.toString()));
        String elementText = this.getTextFromElement(webElement);
        logger.trace(String.format("Got text '%s' from element '%s'", elementText, webElement.toString()));

        return Integer.parseInt(
                CharMatcher.inRange('0', '9').retainFrom(elementText));
    }

    protected List<WebElement> getElements(String elementsXpath) {
        logger.trace(String.format("Get list of WebElements from '%s'", elementsXpath));

        return this.webDriver.findElements(By.xpath(elementsXpath));
    }

    protected WebElement getSubElement(WebElement webElement, String subElementXpath) {
        logger.trace(
                String.format("Get subelement '%s' from the webElement '%s'", subElementXpath, webElement.toString()));

        try {
            WebElement subElement = webElement.findElement(By.xpath(subElementXpath));
            logger.trace(String.format("Sub element found: '%s' with text: '%s'", subElement.toString(),
                    subElement.getText()));

            return subElement;
        } catch (Exception e) {
            return null;
        }
    }

    protected WebElement getSubElement(WebElement webElement, String subElementXpath, int elementIndex) {
        logger.trace(
                String.format("Get subelement '%s' from the webElement '%s'", subElementXpath, webElement.toString()));

        try {
            List<WebElement> subElement = webElement.findElements(By.xpath(subElementXpath));
            for (WebElement webElement2 : subElement) {
                logger.trace(String.format("Sub element found: '%s' with text: '%s'", webElement2.toString(),
                        webElement2.getText()));
            }
            return subElement.get(elementIndex);
        } catch (Exception e) {
            return null;
        }
    }

    protected String getAttribute(String elementXpath, String attributeName) {
        logger.trace(String.format("Get attribute '%s' from element '%s'", attributeName, elementXpath));
        String attributeString = this.webDriver.findElement(By.xpath(elementXpath)).getAttribute(attributeName);

        if (attributeString != null) {
            return attributeString;
        }
        return "";
    }

    protected void selectFromDropdown(String dropdownElementXpath, String elementTextForSelect) {
        logger.trace(String.format("Select from dropdown '%s' element with text '%s'", dropdownElementXpath,
                elementTextForSelect));
        Select dropdownSelect = new Select(this.webDriver.findElement(By.xpath(dropdownElementXpath)));
        dropdownSelect.selectByVisibleText(elementTextForSelect);
    }

    protected void refreshPage() {
        logger.debug("Refresh page");
        this.webDriver.navigate().refresh();
        this.waitPage();
    }

    /**
     * Create WebDriver and login
     */
    protected void startBot() {
        this.webDriver = new ChromeDriver(this.setChromeOptions(BotBase.GOOGLE_PARAMETERS));
        this.login();
    }

    protected void waitPage() {
        try {
            new WebDriverWait(this.webDriver, Duration.ofSeconds(LOGIN_TIMEOUT_SEC))
                    .until(ExpectedConditions.presenceOfElementLocated(
                            By.xpath(APIXpath.xpathElementLoadingOverlay)));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private final void login() {
        logger.info(String.format("Login to '%s'", this.baseURL));
        this.webDriver.manage().deleteAllCookies();
        this.isLoggedIn = false;

        if (this.loginAttempts > BotBase.MAX_LOGIN_ATTEMPTS) {
            long timeDeltaSec = ChronoUnit.SECONDS.between(loginLastAttemptTimestamp, Instant.now());

            if (timeDeltaSec < BotBase.MAX_TIME_DELTA_SEC) {
                logger.error(String.format("Maximum (%d) login attempts reached", BotBase.MAX_LOGIN_ATTEMPTS));
                return;
            } else {
                this.loginAttempts = 0;
            }
        }

        this.loginAttempts++;
        this.loginLastAttemptTimestamp = Instant.now();
        // this.webDriver.manage().window().maximize();
        this.webDriver.get(this.baseURL);
        this.clickButton(APIXpath.xpathButtonLogin);
        this.typeTextInField(APIXpath.xpathTextFieldUsername, this.login);
        this.typeTextInField(APIXpath.xpathTextFieldPassword, this.password);
        this.clickButton(APIXpath.xpathCheckboxRememberMe);
        this.clickButton(APIXpath.xpathButtonAuth);

        this.waitPage();
        this.isLoggedIn = true;
    }

    protected boolean isLoggedIn() {
        return this.isLoggedIn;
    }

    protected void quit() {
        this.webDriver.close();
        this.webDriver.quit();
    }
}
