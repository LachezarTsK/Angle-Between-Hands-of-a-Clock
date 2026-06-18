
package main
import "math"

const MINUTES_IN_HOUR = 60
const HOURS_FULL_CIRCLE = 12
const ANGLE_DEGREES_FULL_CIRCLE = 360

const ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR
const MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE

func angleClock(hour int, minutes int) float64 {
    /*
    minutesOnTheDialForInputHour
    Note: not to be confused with minutes contained in the input hour
    */
    minutesOnTheDialForInputHour := MINUTES_SEPARATING_HOURS_ON_THE_DIAL *
        ((float64(minutes) / float64(MINUTES_IN_HOUR)) + float64(hour%HOURS_FULL_CIRCLE))

    angleDegrees := ANGLE_DEGREES_PER_MINUTE * math.Abs(float64(minutes) - minutesOnTheDialForInputHour)

    return min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees)
}
