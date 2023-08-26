package com.ashokhin.am4bot.model;

public final class MarketingCompany {
    private MarketingCompanyType type;
    private int price;
    private boolean active;

    public MarketingCompany(MarketingCompanyType companyType) {
        this.type = companyType;
    }

    public MarketingCompany(MarketingCompanyType companyType, boolean companyActive) {
        this.type = companyType;
        this.active = companyActive;
    }

    public final MarketingCompanyType getType() {
        return this.type;
    }

    public final String getName() {
        return this.type.getName();
    }

    public final int getPrice() {
        return this.price;
    }

    public final void setPrice(int price) {
        this.price = price;
    }

    public final boolean isActive() {
        return this.active;
    }

    public final void setActive(boolean active) {
        this.active = active;
    }

    public final String getRowXpath() {
        return this.type.getRowXpath();
    }

    public final String getButtonXpath() {
        return this.type.getButtonXpath();
    }

    @Override
    public final String toString() {
        return String.format("%s{type='%s', price='%s', active='%s'}", this.getClass().getSimpleName(), this.type,
                this.price);
    }
}
