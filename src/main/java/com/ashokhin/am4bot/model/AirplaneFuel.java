package com.ashokhin.am4bot.model;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public final class AirplaneFuel {
    private static final Logger logger = LogManager.getLogger(AirplaneFuel.class);
    private static int criticalLevelPercent;
    private FuelType type;
    private int price;
    private int goodPrice;
    private int budgetPercent;
    private int currentCapacity;
    private int maximumCapacity;
    private int holdingCapacity;
    private String displayUnit;

    public AirplaneFuel(
            FuelType type,
            int price,
            int goodPrice,
            int budgetPercent,
            int currentCapacity,
            int maximumCapacity) {
        this.type = type;
        this.price = price;
        this.goodPrice = goodPrice;
        this.budgetPercent = budgetPercent;
        this.currentCapacity = currentCapacity;
        this.maximumCapacity = maximumCapacity;
        this.displayUnit = this.type.getUnit();
        this.updateHoldingCapacity();
    }

    private final void updateHoldingCapacity() {
        this.holdingCapacity = (this.maximumCapacity - this.currentCapacity);
    }

    public final String getType() {
        return this.type.getTitle();
    }

    public final int getPrice() {
        return this.price;
    }

    public static final void setCriticalLevelPercent(int criticalFuelLevelPercent) {
        AirplaneFuel.criticalLevelPercent = criticalFuelLevelPercent;
    }

    public final String getInfo() {
        return String.format(
                "Fuel type: %s\nFuel price: $%d\nFuel good price: $%d\nFuel budget percent: %d%%\nCurrent capacity: %d %s\nMaximum capacity: %d %s\nHolding capacity: %d %s (%d%%)",
                this.type.getTitle(), this.price, this.goodPrice, this.budgetPercent,
                this.currentCapacity, this.displayUnit,
                this.maximumCapacity, this.displayUnit,
                this.holdingCapacity, this.displayUnit,
                (Math.round(this.maximumCapacity / this.holdingCapacity) * 100));
    }

    public final boolean isFull() {
        return (this.currentCapacity <= 0);
    }

    public final boolean notEnoughFuel() {
        return (this.holdingCapacity < (this.maximumCapacity * (AirplaneFuel.criticalLevelPercent * 0.01)));
    }

    public final int getNeedAmount(int accountMoney) {
        if (this.price > this.goodPrice) {
            logger.warn(String.format("'%s' price is too high. Current: $%d, recommended: $%d", this.type.getTitle(),
                    this.price,
                    this.goodPrice));
            if (!this.notEnoughFuel()) {
                return 0;
            } else {
                logger.warn(String.format("Critical '%s' level (less than %d%%). Buy for current price: $%d",
                        this.type.getTitle(), AirplaneFuel.criticalLevelPercent,
                        this.price));
            }
        }

        int availableMoney = (int) Math.round((accountMoney * (this.budgetPercent * 0.01)));
        int maxAvailableCapacity = Math.round((availableMoney / this.price) * 1000);
        int needAmount = (maxAvailableCapacity > this.currentCapacity) ? this.currentCapacity : maxAvailableCapacity;
        int fuelTotalPrice = Math.round((this.price * (needAmount / 1000)));
        logger.debug(String.format(
                "'%s' available money: $%d,\ntotal price: $%d,\nneed amount: %d %s,\nmaximum available amount: %d %s",
                this.type.getTitle(),
                availableMoney, fuelTotalPrice, needAmount, this.displayUnit, maxAvailableCapacity,
                this.displayUnit));

        return needAmount;
    }

    public final void update(int fuelPrice, int currentCapacity) {
        this.price = fuelPrice;
        this.currentCapacity = currentCapacity;
        this.updateHoldingCapacity();
    }

    public final void buyFuelAmount(int needFuelAmount) {
        this.currentCapacity = (this.currentCapacity - needFuelAmount);
        this.updateHoldingCapacity();
    }

    public int getHoldingCapacity() {
        return this.holdingCapacity;
    }

    public int getMaximumCapacity() {
        return this.maximumCapacity;
    }

    @Override
    public final String toString() {
        return String.format(
                "%s{type='%s', price=%d, goodPrice=%d, budgetPercent=%d, currentCapacity=%d, maximumCapacity=%d, holdingCapacity=%d, displayUnit='%s'}",
                this.getClass().getSimpleName(), this.type.getTitle(), this.price, this.goodPrice, this.budgetPercent,
                this.currentCapacity, this.maximumCapacity, this.holdingCapacity, this.displayUnit);
    }
}
