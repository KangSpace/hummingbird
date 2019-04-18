/*
 date: 2019-04-18 22:35:00
 A child is playing with a ball on the nth floor of a tall building. The height of this floor, h, is known.

He drops the ball out of the window. The ball bounces (for example), to two-thirds of its height (a bounce of 0.66).

His mother looks out of a window 1.5 meters from the ground.

How many times will the mother see the ball pass in front of her window (including when it's falling and bouncing?

Three conditions must be met for a valid experiment:
** Float parameter "h" in meters must be greater than 0**
** Float parameter "bounce" must be greater than 0 and less than 1**
** Float parameter "window" must be less than h.**
If all three conditions above are fulfilled, return a positive integer, otherwise return -1.

Note: The ball can only be seen if the height of the rebounding ball is stricty greater than the window parameter.

Example:

h = 3, bounce = 0.66, window = 1.5, result is 3

h = 3, bounce = 1, window = 1.5, result is -1 (Condition 2) not fulfilled).
 */
package main

import "fmt"

func BouncingBall1(h, bounce, window float64) int {
	fmt.Println(h,",",bounce,",",window)
	if h <=0 || !(bounce>0 && bounce < 1) || window >= h{
		return -1;
	}
	count := -1
	for ;h > window; h *= bounce{
		count+=2
	}
	return count
}

func BouncingBall(h, bounce, window float64) int {
	if h <=0 || !(bounce>0 && bounce < 1) || window >=h{
		return -1;
	}
	return 2 + BouncingBall(h*bounce,bounce,window)
}

func main() {
	// 3
	fmt.Println(BouncingBall(3, 0.66, 1.5))
	////3
	//fmt.Println(BouncingBall(40, 0.4, 10))
	////-1
	//fmt.Println(BouncingBall(5, -1, 1.5))
	//15
	fmt.Println(BouncingBall(30 , 0.66 , 1.5))
	//17
	//fmt.Println(BouncingBall(28 , 0.75 , 2.8))
}
