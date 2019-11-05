package yacht

int rollDice(void);
int playGame(void);

int rollDice(void) {
return ((rand() % 6) + 1);
}

int playGame(void){
int dice_1 = 0;
int dice_2 = 0;
int sum = 0;
time_t t;
srand(time(&t));
printf("ROLL THE DICE WITH [ENTER]\n");
dice_1 = rollDice();
dice_2 = rollDice();
sum = dice_1 + dice_2;
if (sum == 7 || sum == 11){
printf("Congratulations you roll %d and WON at your first try!", sum);
}
else {
printf("Your roll was %d ,you lose try agian.\n", sum);