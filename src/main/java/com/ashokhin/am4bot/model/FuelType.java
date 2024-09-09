package com.ashokhin.am4bot.model;

public enum FuelType {
    FUEL("fuel", "Lbs"),
    CO2("co2", "Quotas");

    private String title;
    private String unit;

    private FuelType(String title, String unit) {
        this.title = title;
        this.unit = unit;
    }

    public final String getTitle() {
        return this.title;
    }

    public final String getUnit() {
        return this.unit;
    }

    @Override
    public final String toString() {
        return String.format("%s{title='%s', unit='%s'}", this.getClass().getSimpleName(), this.title, this.unit);
    }
}
