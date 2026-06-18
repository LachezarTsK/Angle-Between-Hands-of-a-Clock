
/**
 * @param {number} hour
 * @param {number} minutes
 * @return {number}
 */
var angleClock = function (hour, minutes) {
    const MINUTES_IN_HOUR = 60;
    const HOURS_FULL_CIRCLE = 12;
    const ANGLE_DEGREES_FULL_CIRCLE = 360;

    const ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR;
    const MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE;

    /*
     minutesOnTheDialForInputHour 
     Note: not to be confused with minutes contained in the input hour
     */
    const minutesOnTheDialForInputHour
            = MINUTES_SEPARATING_HOURS_ON_THE_DIAL * (((minutes / MINUTES_IN_HOUR) + (hour % HOURS_FULL_CIRCLE)));

    const angleDegrees = ANGLE_DEGREES_PER_MINUTE * Math.abs(minutes - minutesOnTheDialForInputHour);

    return Math.min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees);
};
