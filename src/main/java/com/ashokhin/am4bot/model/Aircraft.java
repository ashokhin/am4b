package com.ashokhin.am4bot.model;

public final class Aircraft {
    public static final String TYPE = "data-type";
    public static final String REG_NUMBER = "data-reg";
    public static final String A_CHECK = "data-hours";
    public static final String WEAR = "data-wear";
    private String type;
    private String regNumber;
    private int aCheckHours;
    private int wearPercent;

    public Aircraft(
            String type,
            String regNumber,
            int aCheckHours,
            int wearPercent) {
        this.type = type;
        this.regNumber = regNumber;
        this.aCheckHours = aCheckHours;
        this.wearPercent = wearPercent;
    }

    public final String getType() {
        return this.type;
    }

    public final String getRegNumber() {
        return this.regNumber;
    }

    public final int getACheckHours() {
        return this.aCheckHours;
    }

    public final int getWearPercent() {
        return this.wearPercent;
    }

    @Override
    public final String toString() {
        return String.format(
                "%s{type='%s', regNumber='%s', aCheckHours=%d, wearPercent=%d}",
                this.getClass().getSimpleName(), this.type,
                this.regNumber, this.aCheckHours, this.wearPercent);
    }
}
