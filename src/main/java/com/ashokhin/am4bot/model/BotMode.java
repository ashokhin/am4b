package com.ashokhin.am4bot.model;

public enum BotMode {
    ALL("ALL"),
    REPAIR_LOUNGE("REPAIR_LOUNGE"),
    UPDATE_STUFF_MORALE("UPDATE_STUFF_MORALE"),
    BUY_FUEL("BUY_FUEL"),
    MAINTENANCE("MAINTENANCE"),
    DEPART("DEPART"),
    MARKETING("MARKETING");

    private String title;

    private BotMode(String title) {
        this.title = title;
    }

    public final String getTitle() {
        return this.title;
    }
}
