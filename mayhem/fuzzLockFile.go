package fuzzLockFile

import "strconv"
import "github.com/tfverch/tfvc/internal/lockfile"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {

        case 0:
            content := string(bytes)
            lockfile.ParseProviderSource(content)
            return 0

        case 1:
            content := string(bytes)
            lockfile.ParseProviderPart(content)
            return 0

        default:
            content := string(bytes)
            lockfile.LoadLocks(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}