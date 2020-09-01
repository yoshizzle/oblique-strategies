package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    strategies := []string {
        "Remove specifics and convert to ambiguities",
        "Think of the radio",
        "Don't be frightened of cliches",
        "Allow an easement\n(an easement is the abandonment of a structure)",
        "What is the reality of the situation?",
        "Simple subtraction",
        "Are there sections? Consider transitions",
        "Turn it upside down",
        "Go slowly all the way round the outside",
        "A line has two sides",
        "Infintesimal gradations",
        "Make an exhaustive list of everything you might do and do the last thing on the list",
        "Change instrument roles",
        "Into the impossible",
        "Accretion",
        "Ask people to work against their better judgment",
        "Disconnet from desire",
        "Take away the elements in order of apparent non-importance",
        "Emphasize repitions",
        "Don't be afraid of things because they're easy to do",
        "Is there something missing?",
        "Don't be frightened to display your talents",
        "User unqualified people",
        "Breathe more deeply",
        "Emphasize differences",
        "Only one element of each kind",
        "Do nothing for as long as possible",
        "Bridges\nbuild\nburn",
        "Water",
        "You don't have to be ashamed of using your own ideas",
        "Make a sudden, destructive unpredictable action; incorporate",
        "Tidy up",
        "Use an unacceptable color",
        "Use filters",
        "Balance the consistency principle with the inconsistency principle",
        "Fill every beat with something",
        "Discard an axiom",
        "Listen to the quiet voice",
        "What wouldn't you do?",
        "Is it finished?",
        "Decorate, decorate",
        "Put in earplugs",
        "Give the game away",
        "Reverse",
        "Abandon normal instruments",
        "Trust in the you of now",
        "Use fewer notes",
        "Distorting time",
        "Give way to your worst impulse",
        "Make a blank valuable by putting it in an exquisite frame",
        "The inconsistency principle",
        "Ghost echoes",
        "Don't break the silence",
        "You can only make one dot at a time",
        "Discover the recipes you are using and abandon them",
        "Just carry on",
        "Cascades",
        "(Organic) machinery",
        "Courage!",
        "What mistakes did you make last time?",
        "You are an engineer",
        "Consider different fading systems",
        "Remove ambiquities and convert to specifics",
        "Mute and continue",
        "Look at the order in which you do things",
        "It is quite possible (after all)",
        "Go outside. Shut the door.",
        "Don't stress one thing more than another",
        "Do we need holes?",
        "Cluster analysis",
        "Work at a different speed",
        "Do something boring",
        "Look closely at the most embarrassing details and amplify them",
        "Define an area as 'safe' and use it as an anchor",
        "Mechanicalize something idiosyncratic",
        "Overtly resist change",
        "Emphasize the flaws",
        "Accept advice",
        "Remember those quiet evenings",
        "Take a break",
        "The tape is now the music",
        "Short circuit\n(example; a man eating peas with the idea that they will improve his virility shovels them straight into his lap)",
        "Imagine the music as a moving chain or caterpillar",
        "Use an old idea",
        "Imagine the music as a set of disconnected events",
        "Change nothing and continue with immaculate consistency",
        "Imagine the piece as a set of disconnected events",
        "What are you really thinking about just now? Incorporate.",
        "Children's voices\n-speaking\n-singing",
        "Assemble some of the instruments in a group and treat the group",
        "Feedback recordings into an acoustic situation",
        "Shut the door and listen from outside",
        "Towards the insignificant",
        "Is the tuning appropriate?",
        "Simply a matter of work",
        "Look at a very small object, look at its centre",
        "Not building a wall but making a brick",
        "Revaluation (a warm feeling)",
        "Disciplined self-indulgence",
        "The most important thing is the thing most easily forgotten",
        "Always first steps",
        "Idiot glee",
        "Question the heroic approach",
        "Be extravagant",
        "Always give yourself credit for having more than personality",
        "State the problem in words as clearly as possible",
        "Faced with a choice, do both",
        "Tape your mouth",
        "Twist the spine",
        "Get your neck massaged",
        "Lowest common denominator check\n-single beat\n-single note\n-single riff",
        "Do the washing up",
        "Listen in total darkness, or in a very large room, very quietly",
        "Convert a melodic element into a rhythmic element",
        "Would anybody want it?",
        "Spectrum analysis",
        "Retrace your steps",
        "Go to an extreme, move back to a more comfortable place",
        "Once the search is in progress, something will be found",
        "Only a part, not the whole",
        "Be less critical more often",
    }

    fmt.Println(strategies[rand.Intn(len(strategies))])
}
