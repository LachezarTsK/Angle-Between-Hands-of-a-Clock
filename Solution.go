
import kotlin.math.abs
import kotlin.math.min

class Solution {

    private companion object {

        const val MINUTES_IN_HOUR = 60
        const val HOURS_FULL_CIRCLE = 12
        const val ANGLE_DEGREES_FULL_CIRCLE = 360

        const val ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR
        const val MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE
    }

    fun angleClock(hour: Int, minutes: Int): Double {
        /*
        minutesOnTheDialForInputHour
        Note: not to be confused with minutes contained in the input hour
        */
        val minutesOnTheDialForInputHour =
            MINUTES_SEPARATING_HOURS_ON_THE_DIAL * ((minutes.toDouble() / MINUTES_IN_HOUR) + (hour % HOURS_FULL_CIRCLE))

        val angleDegrees = ANGLE_DEGREES_PER_MINUTE * abs(minutes - minutesOnTheDialForInputHour)

        return min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees)
    }
}
