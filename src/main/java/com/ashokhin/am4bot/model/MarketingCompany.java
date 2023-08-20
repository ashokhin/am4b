package com.ashokhin.am4bot.model;

public class MarketingCompany {
    private String name;
    private int price;
    private boolean active;

    public MarketingCompany(String companyName) {
        this.name = companyName;
    }

    public MarketingCompany(String companyName, boolean companyActive) {
        this.name = companyName;
        this.active = companyActive;
    }

    public final String getName() {
        return this.name;
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
}
