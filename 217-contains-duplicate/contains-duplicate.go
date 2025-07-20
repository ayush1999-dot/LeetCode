func containsDuplicate(nums []int) bool {

    duplicate_check := make(map[int]bool)

    for _,v := range nums{
        if duplicate_check[v]{
            return true
        }
        duplicate_check[v]=true
    }
    return false

    
}