package main

import (
        "sort"
)

func main() {
        // [[10,16], [2,8], [1,6], [7,12]]
        points := [][]int{
                []int{
                        10, 16,
                },
                []int{
                        2, 8,
                },
                []int{
                        1, 6,
                },
                []int{
                        7, 12,
                },
        }
        num := findMinArrowShots(points)
        print(num)
}

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start

// ByEnd implements sort.Interface based on the Right Edge of an Ballon.
type ByEnd [][]int

func (a ByEnd) Len() int           { return len(a) }
func (a ByEnd) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a ByEnd) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func findMinArrowShots(points [][]int) int {
        // 1. find the first balloon from left to right
        // 2. if we shoot an arrow to the right edge of the ballon, we can burst all ballons which
        //    left edge is smaller than the first one.
        // 3. this mean, if the second ballon's left edge is greater than the first one's right edge,
        //    we will need an extra arrow.
        // 4. repeat the loop

        if len(points) == 0 {
                return 0
        }

        sort.Sort(ByEnd(points))

        end := points[0][1]
        arrows := 1

        for i := 1; i < len(points); i++ {
                if points[i][0] > end {
                        arrows++
                        end = points[i][1]
                }
        }

        return arrows
}

// @lc code=end