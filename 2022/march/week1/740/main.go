package main
import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
    if len(nums) == 1 {
		return nums[0]
	}

    tmp := make([]int, 10001)
    maxNum := int(0)
    for _, num := range nums {
        if num > maxNum {
			maxNum = num
		}
        tmp[num] += num
    }

    dp := make([]int, maxNum + 1)
    dp[0] = tmp[0]
    dp[1] = tmp[1]
    for i := 2; i < len(dp); i++ {
        dp[i] = max(dp[i-1], dp[i-2] + tmp[i])
    }
    return dp[maxNum]
}

func max(a, b int) int {
    if b > a {
		return b
	}
    return a
}

// Time Limit Exceeded
func ngSolution(nums []int) int {
	result := int(0)
	sort.Ints(nums)
	sumsMap := map[int]int{}
	uniqueNums := []int{}
	for _, num := range nums {
		if sumsMap[num] == 0 {
			uniqueNums = append(uniqueNums, num)
		}
		sumsMap[num] += num
	}

	// fmt.Println(uniqueNums, sumsMap)
	helper(0, uniqueNums, sumsMap, map[int]int{}, &result)
	return result
}

func helper(sum int, nums []int, sumsMap map[int]int, memo map[int]int, result *int) {
	// fmt.Println(sum, nums)
	n := len(nums)
	if n == 0 {
		if sum > *result {
			*result = sum
		}
		return
	}

	target := nums[0]
	if sum < memo[target] {
		return
	}
	memo[target] = sum

	if sumsMap[target + 1] == 0 {
		helper(sum + sumsMap[target], nums[1:], sumsMap, memo, result)
	} else {
		helper(sum, nums[1:], sumsMap, memo, result)
		helper(sum + sumsMap[target], nums[2:], sumsMap, memo, result)
	}
}

func generateNums(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 10000 {
			result[i] = i + 1 - 10000
			continue
		}
		result[i] = i + 1
	}
	return result
}

func main() {
	// result: 6
	// nums := []int{3, 4, 2}

	// result: 9
	// nums := []int{2, 2, 3, 3, 3, 4}

	// result: 1
	// nums := []int{1}

	// result: 30
	// nums := []int{1,2,3,4,5,6,7,8,9,10}

	// result: 
	// nums := generateNums(2000)

	// result: 531
	// nums := []int{1,2,3,4,5,6,7,8,9,10,4,360,7,8,2,9,1,67,4,8,2,3,1,46}

	// result: 4
	// nums := []int{3, 1}

	// result: 138
	// nums := []int{1,8,5,9,6,9,4,1,7,3,3,6,3,3,8,2,6,3,2,2,1,2,9,8,7,1,1,10,6,7,3,9,6,10,5,4,10,1,6,7,4,7,4,1,9,5,1,5,7,5}

	// result: 614204
	nums := []int{404,134,996,652,12,116,401,936,425,935,625,662,278,6,225,680,800,982,345,23,434,251,635,541,687,826,528,130,958,44,676,430,651,401,982,523,48,192,985,245,896,393,772,51,766,667,238,775,956,695,852,956,66,809,705,671,291,892,570,129,631,745,334,973,897,237,308,556,650,416,304,911,910,396,219,953,960,439,749,107,714,321,895,365,864,282,989,219,117,174,904,797,731,64,165,548,430,933,579,630,915,546,705,860,499,880,942,90,995,184,576,903,769,760,586,762,717,394,278,751,570,589,151,781,483,397,918,241,506,410,123,831,779,518,956,621,164,2,45,173,612,947,567,390,460,898,640,649,317,398,324,116,150,316,843,720,780,550,698,597,136,270,475,860,394,418,392,93,710,139,121,519,886,181,154,767,318,294,314,585,442,378,863,881,726,273,199,39,865,906,965,898,97,558,534,70,632,498,686,948,262,731,834,475,847,498,510,180,331,646,992,335,714,68,546,534,753,366,111,270,134,295,33,347,868,661,522,926,601,769,432,268,514,56,675,203,280,128,467,452,707,633,118,770,56,70,828,402,336,961,125,831,803,237,410,51,487,822,26,788,719,563,905,162,420,415,916,252,564,167,73,775,740,795,565,939,286,178,86,450,740,333,769,876,9,262,325,424,787,266,439,929,662,396,127,212,90,308,778,593,149,899,920,279,496,259,213,325,910,63,22,705,440,689,413,462,132,388,20,624,540,892,244,561,892,865,397,258,5,766,243,931,165,479,802,980,271,455,286,664,932,267,447,436,91,428,228,561,840,61,639,301,906,987,201,210,518,461,893,788,116,251,448,347,277,609,465,317,473,479,806,989,957,386,723,555,691,103,289,489,692,42,744,478,648,616,798,720,826,445,692,264,846,725,417,253,596,719,347,41,183,857,119,254,827,714,340,937,85,753,769,757,15,685,760,354,248,792,956,387,140,172,540,591,938,511,410,239,639,552,717,754,494,136,908,3,145,102,609,49,458,221,400,29,245,920,65,216,237,769,788,967,744,999,539,752,455,407,739,872,383,138,757,872,515,377,431,836,684,26,378,744,793,986,664,270,709,837,166,126,512,793,941,796,103,570,106,504,516,477,587,395,448,436,520,4,376,33,737,780,598,179,940,860,464,219,326,598,202,984,654,587,411,174,177,965,634,164,565,546,956,591,52,33,998,242,866,105,48,848,131,569,493,439,806,731,518,588,866,467,266,268,519,478,368,458,528,684,677,198,241,924,980,606,33,52,639,255,340,185,113,685,876,482,576,944,245,673,15,324,949,454,613,516,870,818,549,706,793,215,44,418,56,12,508,943,128,472,172,795,711,836,764,78,757,32,956,405,565,627,159,239,634,665,217,542,779,924,284,722,580,482,392,754,145,710,328,416,378,478,39,611,717,696,961,83,452,97,777,913,694,560,303,920,118,858,695,977,270,174,505,569,424,909,438,630,141,586,762,495,144,785,271,181,930,291,132,63,863,827,929,821,154,637,743,461,889,770,73,88,104,365,992,711,32,823,225,995,376,967,436,843,540,636,22,142,811,116,578,487,505,81,164,432,256,251,421,905,100,745,628,491,844,398,526,34,104,567,570,728,195,736,155,937,224,343,626,869,384,898,640,527,286,727,672,669,785,99,643,108,568,912,307,122,931,721,87,647,338,46,276,985,375,112,471,26,4,170,57,630,703,341,986,175,611,525,168,207,201,689,407,406,684,492,850,914,980,396,857,736,49,590,219,255,104,495,102,26,910,919,614,855,523,681,819,202,951,457,289,658,860,172,849,912,482,658,531,803,269,976,775,529,105,779,915,280,789,686,194,749,378,176,592,151,105,427,37,779,606,903,801,397,536,276,31,494,428,49,818,11,256,421,478,594,289,380,754,261,284,574,256,219,582,68,71,808,520,968,781,879,319,129,569,273,943,602,491,102,449,825,410,848,535,334,547,273,342,390,609,220,10,483,218,793,393,831,272,945,315,8,405,954,424,958,739,347,726,824,236,143,538,648,631,243,768,506,353,204,829,889,683,871,646,294,323,465,291,449,913,841,51,768,68,512,872,819,327,115,688,899,491,521,466,954,298,302,201,379,677,954,263,625,260,365,141,768,433,77,405,549,446,792,538,341,124,123,893,770,161,555,642,98,326,43,178,921,390,366,511,878,622,260,988,339,777,12,376,333,421,298,671,810,101,186,210,625,386,920,209,19,551,864,225,123,51,772,181,547,844,171,883,649,906,7,989,196,232,932,25,715,876,820,87,140,42,940,813,466,518,405,302,267,888,965,574,491,747,987,840,32,645,822,263,482,658,601,151,638,266,546,265,125,870,555,413,461,10,185,976,812,379,877,179,303,480,869,621,381,1,975,115,741,903,970,524,370,292,774,593,335,417,38,533,898,602,56,173,287,399,728,764,546,82,187,452,376,884,882,624,155,580,718,179,834,512,988,15,260,118,594,479,317,508,887,994,303,578,528,939,808,498,203,373,386,696,835,155,620,315,366,823,282,96,215,660,406,835,48,442,895,597,126,51,651,8,688,325,148,572,727,15,992,527,500,745,925,973,605,740,608,180,10,501,147,632,640,682,490,222,852,502,978,111,291,508,140,799,323,471,977,416,748,889,752,364,789,889,1000,367,411,399,188,135,873,587,671,561,615,221,910,48,986,421,775,441,923,547,867,401,566,245,855,117,100,325,438,30,532,32,443,618,125,443,735,908,744,329,345,465,953,9,946,109,566,401,355,150,424,214,508,471,139,134,632,454,634,425,204,400,110,746,975,667,677,51,31,32,438,919,252,279,845,111,68,875,938,521,220,647,463,442,258,309,968,472,209,559,455,995,618,281,583,461,912,561,460,893,738,679,759,662,26,697,657,801,117,384,474,703,410,773,315,966,451,986,921,975,483,258,663,734,470,576,762,400,349,667,766,588,468,380,992,143,717,412,24,595,811,130,283,351,11,97,517,833,13,596,656,242,404,63,908,853,454,242,134,352,196,81,271,622,887,349,39,889,365,770,806,741,867,246,346,470,830,777,965,622,722,603,380,693,531,410,738,18,372,622,505,211,397,315,811,452,990,920,218,631,594,821,702,973,272,633,779,248,5,914,325,150,417,913,28,264,378,117,299,210,400,131,992,290,347,758,468,478,187,814,446,824,662,576,248,841,144,395,918,837,821,885,359,389,441,15,39,413,347,682,17,262,754,123,830,542,965,421,196,372,115,348,986,230,638,316,658,634,495,592,511,920,946,673,549,720,924,262,750,369,562,644,481,987,65,264,636,169,107,478,523,122,518,652,585,831,712,515,912,478,47,194,207,747,456,791,428,176,818,741,818,441,932,991,314,148,395,585,14,438,519,835,334,548,842,31,944,154,369,918,501,119,185,372,940,887,980,149,229,347,801,701,832,967,719,903,78,835,688,288,718,235,72,357,750,74,937,635,75,49,473,890,59,324,630,660,758,153,70,438,326,520,395,205,769,909,254,148,344,744,225,560,294,681,810,36,188,498,325,699,890,127,140,52,260,362,439,33,26,196,798,999,274,822,826,864,660,322,449,917,387,291,923,388,233,980,154,364,634,192,818,406,959,384,801,602,851,262,125,308,354,928,293,790,191,112,967,471,615,108,556,207,98,740,947,420,984,455,169,305,815,61,203,749,679,160,780,192,87,50,391,812,324,143,241,231,282,272,406,595,887,126,885,78,398,411,387,80,23,201,441,807,409,28,181,198,732,8,908,583,780,390,318,863,519,21,253,271,855,403,848,341,817,530,423,470,266,450,41,355,999,285,187,611,208,505,42,512,833,323,725,989,414,34,727,653,398,843,60,131,13,3,530,730,252,901,803,661,95,296,94,178,334,889,769,110,430,588,6,54,382,482,841,239,436,125,610,534,423,753,347,592,451,279,107,988,650,613,10,249,787,916,871,786,646,711,398,41,374,59,649,636,353,841,455,888,479,454,49,921,453,703,622,205,942,706,751,466,881,350,155,254,917,76,741,751,713,723,846,557,295,138,301,629,674,924,10,89,622,559,107,905,36,555,13,833,783,71,695,406,466,356,259,779,24,350,794,715,266,648,142,47,330,753,272,585,665,881,231,825,532,626,448,441,841,903,733,224,54,470,449,92,44,778,334,555,496,876,701,274,543,754,782,688,155,756,867,658,26,377,24,673,777,473,962,498,965,608,72,631,220,357,466,380,625,991,21,112,932,377,813,943,261,832,886,304,206,790,536,548,400,283,905,116,680,918,474,376,373,287,549,570,10,646,836,564,972,489,478,899,263,362,570,720,876,374,16,655,64,521,631,429,224,790,844,213,396,559,943,242,335,225,582,538,480,614,697,240,412,755,125,448,916,78,883,416,233,616,578,311,796,99,332,497,436,210,142,392,73,938,433,252,202,696,592,109,872,906,953,445,113,82,606,520,646,55,795,350,353,884,919,572,477,357,820,526,662,867,868,283,755,323,529,203,700,656,807,757,343,587,658,106,441,941,417,738,885,750,506,403,859,538,655}

	// result: 
	// nums := []int{}

	result := deleteAndEarn(nums)
	fmt.Printf("result = %v\n", result)
}
