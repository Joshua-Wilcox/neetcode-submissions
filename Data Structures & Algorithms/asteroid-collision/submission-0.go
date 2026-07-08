func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0)

    for _, a := range asteroids {
        alive := true
        for alive && a < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
            top := stack[len(stack)-1]
            if top < -a {
                // top asteroid destroyed, keep checking against next one
                stack = stack[:len(stack)-1]
            } else if top == -a {
                // both destroyed
                stack = stack[:len(stack)-1]
                alive = false
            } else {
                // current asteroid destroyed
                alive = false
            }
        }
        if alive {
            stack = append(stack, a)
        }
    }

    return stack
}