package com.ashokhin.am4bot.model;

public enum MarketingCompanyType {
    AIRLINE_REPUTATION(
            "Airline reputation",
            APIXpath.xpathElementFinanceMarketingCompany1,
            APIXpath.xpathButtonFinanceMarketingCompany1Do),
    ECO_FRIENDLY(
            "Eco friendly",
            APIXpath.xpathElementFinanceMarketingCompany2,
            APIXpath.xpathButtonFinanceMarketingCompany2Do);

    private String name;
    private String rowXpath;
    private String buttonXpath;

    private MarketingCompanyType(
            String name,
            String rowXpath,
            String buttonXpath) {
        this.name = name;
        this.rowXpath = rowXpath;
        this.buttonXpath = buttonXpath;
    }

    public final String getName() {
        return this.name;
    }

    public final String getRowXpath() {
        return this.rowXpath;
    }

    public final String getButtonXpath() {
        return this.buttonXpath;
    }

    @Override
    public String toString() {
        return String.format("%s{name='%s', rowXpath='%s', buttonXpath='%s'}", this.getClass().getSimpleName(),
                this.name, this.rowXpath, this.buttonXpath);
    }
}
