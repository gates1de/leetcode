package main
import (
	"fmt"
)

func minOperations(nums []int, x int) int {
	result := int(-1)
	leftSum := int(0)
	rightSum := int(0)
	// left index
	i := int(0)
	// right index
	j := len(nums) - 1
	for i <= j {
		leftCost := i
		rightCost := len(nums) - j - 1
		currentSum := leftSum + rightSum
		// fmt.Printf("i = %v, j = %v, leftCost = %v, rightCost = %v, leftSum = %v, rightSum = %v\n", i, j, leftCost, rightCost, leftSum, rightSum)

		leftmost := nums[i]
		rightmost := nums[j]
		// fmt.Printf("leftmost = %v, rightmost = %v\n", leftmost, rightmost)

		if i == j {
			if currentSum + leftmost == x {
				result = min(result, leftCost + rightCost + 1)
			}
			break
		}
		if leftmost > x && rightmost > x {
			return - 1
		}

		// reach x cases
		if currentSum == x {
			result = min(result, leftCost + rightCost)
		}

		if currentSum + leftmost + rightmost == x {
			result = min(result, leftCost + rightCost + 2)
		} else if leftSum + leftmost == x {
			result = min(result, leftCost + 1)
		} else if rightSum + rightmost == x {
			result = min(result, rightCost + 1)
		} else if currentSum + leftmost == x || currentSum + rightmost == x {
			result = min(result, leftCost + rightCost + 1)
		}

		if leftSum + leftmost < x {
			leftSum += leftmost
		}
		if rightSum + rightmost < x {
			rightSum += rightmost
		}
		if leftmost <= x {
			i++
		}
		if rightmost <= x {
			j--
		}
	}
	return result
}

func min(a, b int) int {
	if a < 0 {
		return b
	} else if b < 0 {
		return a
	} else if a < b {
		return a
	}
	return b
}

func main() {
	// nums := []int{1, 1, 4, 2, 3}
	// x := 5

	// nums := []int{5, 6, 7, 8, 9}
	// x := 4

	// nums := []int{3, 2, 20, 1, 1, 3}
	// x := 10

	// nums := []int{10, 1, 1, 1, 1, 1}
	// x := 5

	// nums := []int{8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309}
	// x := 134365

	// nums := []int{716,9207,8240,4703,7194,3762,4401,291,4902,5062,7580,1470,281,4856,3767,4682,7512,6109,6650,5946,3018,8286,8351,2498,8128,6579,6003,1637,2676,7225,1072,1265,3776,1147,4308,2147,8967,2405,1635,7641,9504,6550,7046,8047,8715,3866,7984,2160,3264,4954,7840,7939,5158,1954,6000,7279,7645,484,9973,9869,9911,9988,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,9978,7338,4628,4687,3989,3590,6866,1498,4659,8360,2249,5315,3446,9412,1011,1330,5364,6844,8616,9241,4363,5442,1700,1787,1879,7563,4795,4670,6095,7327,7002,7374,9571,3416,7286,4980,6780,5440,6263,9365,3747,6872,9332,6494,1367,4727,4840,3430,5616,6319,8785,8762,5321,8943,5526,4732,5352,3780,1209,4019,7544,505,2973,1882,8431,9444,181,3390,7374,8185,1943,7939}
	// x := 1572708

	// nums := []int{1583,400,7633,1321,4473,565,1234,2342,6364,8738,6320,790,3306,1272,8987,8920,3273,8778,6421,616,909,4292,6959,7822,615,7777,590,2368,1638,9811,51,7028,1814,8108,6976,4769,8825,3308,3955,1240,3055,4039,5593,1548,2776,4073,2819,8429,7514,4906,5955,5387,6557,3987,8809,4826,8876,7006,5873,1580,1062,1438,3511,2125,5530,3831,8565,6672,1743,7083,7445,7142,4393,4104,4209,9034,8512,8032,1301,7539,3100,3929,5669,2080,6531,724,7920,6074,1967,9529,5716,4931,2715,9619,2203,108,609,9194,8880,3550,2834,2998,8587,1755,5334,6264,4051,1254,8820,5466,1932,1408,3297,2594,803,4649,401,4985,2687,7489,9640,8528,7637,7878,6453,4027,2449,3682,2582,1109,841,651,2277,7189,7096,2905,6195,8283,5212,6393,1157,1632,9277,3829,9798,5752,2112,511,1723,3423,6851,8083,3857,1719,2632,3494,9344,7924,3433,5147,9489,9404,6001,3456,2828,9259,741,4261,8273,23,6762,8651,8594,7693,5881,2439,6979,9713,6242,7692,6344,7224,9636,5909,8723,8718,4063,1671,2089,8930,3938,7091,8937,7237,8324,3134,4529,569,571,7497,6795,6509,8671,9584,6272,6573,4440,4962,7558,8291,2,3889,5317,1050,1625,5851,6373,4643,9854,8329,7579,4555,2922,1879,1213,4673,6595,658,3221,8614,1420,206,7775,2314,7195,6100,6155,5925,7308,2491,9088,8291,8699,2013,3261,9273,9056,9217,320,4013,7668,7751,5593,7319,5757,2270,7258}
	// x := 128443668

	nums := []int{5281,4109,5962,1122,8456,7256,9646,1708,7754,3164,8637,3769,5311,3714,2661,3807,5270,2026,164,2510,9151,9796,2563,5234,297,2883,5931,8937,9729,5201,9954,6776,5631,2966,1316,3999,2302,8014,7120,4012,9610,6541,7009,3438,9722,4564,9361,5800,5404,9304,1483,882,9834,2592,8812,371,7574,8236,2698,7254,4212,9254,5228,9219,5310,6896,1693,9635,1544,5751,7346,4104,4867,977,4988,5981,1644,7039,5808,8544,8961,2364,6156,6468,6357,4003,5106,261,3490,7859,2876,9116,3955,1218,2598,5866,1613,3737,8047,8483,2450,9009,9806,8432,8006,616,4519,2180,9641,9988,9151,1902,7002,8397,7355,9236,2302,5218,7151,3728,3118,8120,1728,2963,3588,4535,1953,2923,9496,7013,1692,6697,4962,862,3406,3484,9902,3031,6706,4025,6850,5257,7176,6724,1669,6801,9857,8702,1430,8299,3070,5729,3496,9443,6430,5988,7419,4436,9946,7153,6561,9938,6154,8614,1181,9708,2362,7881,9043,1703,5874,1257,5929,5729,606,8534,7245,8213,8227,2170,6374,2483,2111,4013,440,6240,6458,3812,15,7167,7967,1863,3034,742,7945,1412,8673,6499,8897,4716,816,4596,5337,8502,8974,690,9339,2095,4330,7884,8226,6108,8752,8657,1547,7796,4552,1913,3445,9278,5373,7569,1569,4422,2732,2833,5756,7402,1034,8197,8275,5748,8535,6599,2872,9128,3,9739,8336,2451,189,3264,878,504,2119,4518,2771,4082,6175,4135,5465,8331,7940,2665,6176,7136,5674,4746,6348}
	x := 151230

	result := minOperations(nums, x)
	fmt.Printf("result = %v\n", result)
}

