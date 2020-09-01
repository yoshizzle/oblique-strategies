package strategies

import (
    "fmt"
    "math/rand"
    "time"
)

function Strategy(strategy string) string {
    strategy := fmt.Sprintf(randomFormat(), strategy)
    return strategy
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    strategies := []string {
        "strategy 1",
        "stragegy 2",
        "strategy 3",
        "strategy 4",
    }

    return strategies[rand.Intn(len(strategies))]
}
