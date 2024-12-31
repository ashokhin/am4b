package com.ashokhin.am4bot.model;

public final class HubsStats {
    private String name;
    private String type;
    private Float value;

    public HubsStats() {

    }

    public HubsStats(String name, String type, Float value) {
        this.name = name;
        this.type = type;
        this.value = value;
    }

    public final String[] getLabels() {
        return new String[] { this.name, this.type };
    }

    public final Float getValue() {
        return this.value;
    }
}
