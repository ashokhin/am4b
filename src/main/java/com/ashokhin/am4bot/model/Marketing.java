package com.ashokhin.am4bot.model;

import java.util.HashMap;

public class Marketing {
    public static final HashMap<String, HashMap<String, String>> marketingCompaniesMap = new HashMap<String, HashMap<String, String>>() {
        {
            put(
                    "Airline reputation",
                    new HashMap<String, String>() {
                        {
                            put("rowXpath", APIXpath.xpathElementFinanceMarketingCompany1);
                            put("buttonXpath", APIXpath.xpathButtonFinanceMarketingCompany1Do);
                        }
                    });

            put(
                    "Eco friendly",
                    new HashMap<String, String>() {
                        {

                            put("rowXpath", APIXpath.xpathElementFinanceMarketingCompany2);
                            put("buttonXpath", APIXpath.xpathButtonFinanceMarketingCompany2Do);
                        }
                    });
        }
    };

    public static final String MARKETING_COMPANY_REPUTATION_DURATION = "24 Hours";
}
