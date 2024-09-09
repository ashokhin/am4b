package com.ashokhin.am4bot.model;

public enum MaintenanceOperation {
    A_CHECK("A-Check"),
    REPAIR("Repair"),
    MODIFY("Modify");

    private String title;

    private MaintenanceOperation(String title) {
        this.title = title;
    }

    public final String getTitle() {
        return this.title;
    }

    @Override
    public final String toString() {
        return String.format("%s{title='%s'}", this.getClass().getSimpleName(), this.title);
    }
}
