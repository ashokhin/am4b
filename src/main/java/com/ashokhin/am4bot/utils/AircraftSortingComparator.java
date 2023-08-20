package com.ashokhin.am4bot.utils;

import java.util.Comparator;

import com.ashokhin.am4bot.model.Aircraft;

public final class AircraftSortingComparator implements Comparator<Aircraft> {

    @Override
    public int compare(Aircraft o1, Aircraft o2) {
        return o1.getRegNumber().compareTo(o2.getRegNumber());
    }

}
