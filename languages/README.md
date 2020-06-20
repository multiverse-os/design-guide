# C Experiments 
Software engineering should be treated as a science, not a trade. When we fail
to recognize that computer science is the heart of good software enginering, we
have planes literally fall out of the sky, as we saw with Boeing poor accessment
on the importance of good software engineering, an intelligent and properly
designed engineering lifecycle that met the requirements of the project. 

#### A compuer lab
One of the beautiful things about computer science, is that unlike many other
fields of science, experiments are far simpler to setup, conduct, and even
replicate. And a good indicator on weather you are approaching software
engineering as a computer scientists or approaching programming as a trade is:
are you regularly conducting experiments that follow the scientific method?
Even if these are experiments that have arleady been proven by others, it is
important for scientists to regularly challenge their intuition, their held
beliefs, and what they are told through experimentation, replicated across many
different environments, as many times as reasonably possible. 

What that means, will depend on the complexity of your computer lab, and if you
are saying to yourself right now: "But I don't have a compuer lab..."; well you
have just assigned yourself your first task. Especially when taking on a task
large as operating system development, a developer needs a computer lab, or a
variety of computers networked together, preferably by professional routers,
which can be anything from an old Cisco router you bought off ebay, or the FPGA
you setup to emulate a few Cisco router models, or a linux computer with as many
network interface cards (NICs) as you can find, and setup to te dedicated task
of routing your computer lab. A common design is two routers, providing two
segregated networks.
 
If you can not afford actual hardware, do not feel like you are excluded, this 
is where software like `portal-gun` steps in; as it can easily setup a computer 
lab environmental using virtual machines. 

Even two physical machines with a virtual router and virtual machines on each
connected together function as a nice minimal setup; but everything can be
virtual if necessary. 

#### I setup a lab, now what? 
Well now you can start testing proof-of-conceps (POCs), from security related,
to preformance related. You can setup complex networks, and test for leaks, and
now with fuzzing, you can attack with random data for whatever amount of time
you can afford to invest; and see if there are any edge cases that would have
been incredibly difficult to catch with tests or manual input. 

A computer lab is what a greenhouse is to the plant scientist, you need a place
to setup controlled environments, to run experiments and disprove hypothesis. 

If you are unfamiliar, lets just review the concept: 

> "A hypothesis is the cornerstone of the scientific method. It is an educated 
> guess about how the world works that integrates knowledge with observation. 
> Everyone appreciates that a hypothesis must be testable to have any value, 
> but there is a much stronger requirement that a hypothesis must meet. A 
> hypothesis is considered scientific only if there is the possibility to 
> disprove the hypothesis." - Internet

A hypothesis or model is called falsifiable if it is possible to conceive of an
experimental observation that disproves the idea in question. 

One of the wonderful things about computer science is that it is very easy to
conduct an experiment, especially once you have taken the time to setup a
computer lab. 

Unlike a scientist studying trees, where a hypothesis may take literally decades
to falisify, we can build scripts and setup an experiment and test it within
minues, and often even long experiments will not take more than hours, days or
weeks. 

And while replication is a massive problem to all sciences currently, especially
with so many black-box pieces of scientific equipment, that prevent scientists
not from just learning the basics of what they are doing, but literally make it
impossible for them to tell if a problem in their experiment is their design,
some outside factor, or a firmware, or software bug that they are unable to peer
review, unable to fix, and since different hardware uses different firmware and
software, often experiments simply can't be conducted across labs with any
reasonable expectation that actual experimental replication. 

Computer science, thanks to the our hippy forefathers, like Richard Stallman,
who like everyone I know, is not perfect, has made it his life's mission to
ensure software, function like math, and be open and shared between scientists.
And truly our current open source methodologies, ideaology, and development
tools, do not go far enough. But thanks to them, we are can still work knowing
every layer of the stack, isolate exactly where an issue is coming from and
truly conduct science. A true rarity, when compared to the closed source
machines that cost often 1,000x their worth, lack basic functionality, and are
guaranteed to be riddled with firmware and software bugs that no one will find
since their is no peer review, so no one will be forced to fix them, or admit
they exist; resulting in not just holding back science, but also while gouging 
scientific insutitions for shoddy equipment while doing it.

So be thankful, you can actually practice science, and do it, setup a lab, and 
don't make assumptions, don't take people at their word: hypothesize, test,
experiment, and share your results, methods, and tools needed to replicate. 

Multiverse OS developers seek to set a standard for replication, and hope that
portal-gun can not just provide the backbone for secure computing in Multiverse
OS, but also provide the structure for reliable experimentation, experimental
replication, and mutually beneficial experimentation. 

#### Gray Literature, and a New Way
Multiverse OS developers decided early one to launch a network theory gray
literature publication with open submissions and free journals accessible by
anyone intersted. 

To do this, our plan is to build the tools for operating a gray literature peer
reviewed publication; but not stop there, to change the expectations of peer
reviewed literature, and to introduce new concepts and requirements. For
example, we want to treat articles as ongoing discussions, with more than two
peer reviews, that branch and fork, into a graph of experiments, articles, and
related experiments. This approach better reflects the technology we use, over a
model that is built around the limitations of the time it was first introduced:
the scientific article, a flag, non-interactive article, made sense for when
they were first being produced, but in 2020, they no longer make sense. 

All scientific articles, should include their entire dataset, all their
equations, all the math, and all the source code. This would allow others to
work with the data and ensure there is not "playing with the numbers"; but it
would also allow the work acheived to be repblicated, across as many different
environments as possible, and allow the results to be used in a variety of
different methods and applications. Using the modern \*.epub format, we could
introduce new interactive figures beyond simple graphs, limited to 2D, static
representations. We could include full consoles, for interacting with data,
every graph could flip around to show the data table used to generate it. And
every "article" could include everything necessary not just replicate it, but to
replicate it and submit signed results to be included into the ongoing
discussion. 

As replications come in from around the world, the peer-reviews could also come
in, not just reviewing the article in its entirity, but could comment,
line-by-line, by paragraph, by sections of code, that could be included in the
review, and tracked by time, and replies, could form a discussion with context. 

This would also do away with the concept of # of citations = better paper. Which
is a naive and short sighted determination of what is useful, good, or even
beneficial to scientific discussion. These sort of easily game-able metrics, are
the domain of hackers, and getting hackers involved will immediately expose
these sorts of games for what they are: petty, useless, and meaningless metrics
used by organizations to legitimize themselves, for funding or prestige; rightly
or wrongly, but regardless the reality. 


#### Where are we now?
If Miltiary science is going to continue to outpace everything, as outfits like
the NSA choose to reveal only some of their findings, so they can have working
0-days, or submit backdoored projects publically, only sometimes making real co
ntributions. And the corporate environment, turns everything, including public
and open work into black boxes for profit at the cost of scientific progress;
and both saboging the limited sphere of academia, so that academia works for
their profit, or benefit; then we need to secure a new, fourth path for science;
some call it citizen science, but it will need to be bigger than that, it will
require funding, it will require p2p networks, peer-review, cryptographic based
identities to verify historical contributions, and secure real funding through a
variety of means including but not limited to donations, subscriptions, but also
open source workers cooperatives, taking government contracts in small groups,
and pooling together our collective resources, for a collective ownership of
both ideas and property. 

And this all starts, with you building that computer lab, so get started. 
