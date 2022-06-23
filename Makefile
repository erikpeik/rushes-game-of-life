NAME = life

DGRAY = \033[1;30m
NCOLOR = \033[0m
PURP = \033[1;35m
WHITE = \033[1;37m
GREEN = \033[1;32m
BLUE = \033[1;34m

all: $(NAME)

$(NAME):
	@clear
	@echo "$(BLUE)game_of_live v.0.1"
	@echo "$(PURP)by: vkuikka & emende"
	@echo "$(DGRAY)"
	@echo "Building $(NAME)...$(WHITE)"
	go build -o life .
	@echo "$(DGRAY)Building $(SPLIT_2)...$(WHITE)"
	go build -o life_animated other_versions/animated/life_animated.go
	@echo "$(DGRAY)Building $(SPLIT_4)...$(WHITE)"
	go build -o life_split2 other_versions/split2/life_split2.go
	@echo "$(DGRAY)Building $(ANIMATED)...$(WHITE)"
	go build -o life_split4 other_versions/split4/life_split4.go
	@echo ""
	@echo "$(GREEN)All good! :)"
	@echo ""
	@echo "$(BLUE)Usage: ./life initial_state iterations"
	@echo "$(NOCOLOR)"

clean:
	rm -f life
	rm -f life_animated
	rm -f life_split2
	rm -f life_split4

re: clean all
