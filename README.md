Ссылка на первоисточник задания: https://docs.google.com/document/u/0/d/1Dv_6AIhxg_3ezu6VMcEnMpyfRzgym9l8PmE4ULGfjgM/mobilebasic

Примечания:
В Digits в Naturals цифры идут слева направо, то есть нулевой элемент - старший разряд

Naturals - натуральные с нулём

# Распределение заданий

Номер | Натуральные числа с нулем:(n; A[..]) - номер старшей позиции и массив цифрЦифра D (тип - целое) | Имена | Базовые модули, на которые ссылается данный модуль | Человек, ответственный за модуль
-- | -- | -- | -- | --
N-1 | Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе. | Compare (COM_NN_D) | | Турбина
N-2 | Проверка на ноль: если число не равно нулю, то «да» иначе «нет» | CheckNull (NZER_N_B) | | Турбина
N-3 | Добавление 1 к натуральному числу | Addition1 (ADD_1N_N) | | Хвостовский
N-4 | Сложение натуральных чисел | Addition (ADD_NN_N) | Compare | Семёнов
N-5 | Вычитание из первого большего натурального числа второго меньшего или равного | Subtraction (SUB_NN_N) | Compare | Лицеванова
N-6 | Умножение натурального числа на цифру | MultiplicationNaturalNumber (MUL_ND_N) | | Хвостовский
N-7 | Умножение натурального числа на 10^k | MultiplicationBy10k (MUL_Nk_N) | | Хвостовский
N-8 | Умножение натуральных чисел | Multiplication (MUL_NN_N) | MultiplicationNaturalNumber MultiplicationBy10k Addition | Грунская
N-9 | Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом | DifferenceOfNaturals (SUB_NDN_N) | Subtraction MultiplicationNaturalNumber Compare  | Комаровский
N-10 | Вычисление первой цифры деления большего натурального на меньшее, домноженное на 10^k,где k - номер позиции этой цифры (номер считается с нуля) | DivideOneIteration (DIV_NN_Dk) | MultiplicationBy10k Compare | Тростин
N-11 | Частное от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля) | IntegerFromDivision (DIV_NN_N) | DivideOneIteration DifferenceOfNaturals | Пименов
N-12 | Остаток от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля) | RemainderFromDivision (MOD_NN_N) | IntegerFromDivision DifferenceOfNaturals | Пименов
N-13 | НОД натуральных чисел | GreatestCommonDivisor (GCF_NN_N) | RemainderFromDivision Compare CheckNull| Турбина
N-14 | НОК натуральных чисел | LeastCommonMultiple (LCM_NN_N) | GreatestCommonDivisor Multiplication | Турбина

Номер | Целые числа (b, n; A[..]) - знак числа (1 — минус, 0 — плюс) номер старшей позиции и массив цифр |   |   | Человек, ответственный за модуль
-- | -- | -- | -- | --
Z-1 | Абсолютная величина числа, результат - натуральное | Absolute (ABS_Z_N) | | Тростин
Z-2 | Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное) | Positivity (POZ_Z_D) | | Турбина
Z-3 | Умножение целого на (-1) | MultiplicationByNegativeOne (MUL_ZM_Z) | | Хвостовский
Z-4 | Преобразование натурального в целое | FromNaturalsToWhole (TRANS_N_Z) | | Лицеванова
Z-5 | Преобразование целого неотрицательного в натуральное | FromWholeToNaturals (TRANS_Z_N) | | Лицеванова
Z-6 | Сложение целых чисел | Addition (ADD_ZZ_Z) | Positivity Absolute Compare Addition Subtraction MultiplicationByNegativeOne | Семёнов
Z-7 | Вычитание целых чисел | Subtraction (SUB_ZZ_Z) | Positivity Absolute Compare Addition Subtraction MultiplicationByNegativeOne | Семёнов
Z-8 | Умножение целых чисел | Multiplication (MUL_ZZ_Z) | Positivity Absolute Multiplication MultiplicationByNegativeOne | Тростин
Z-9 | Частное от деления целого на целое (делитель отличен от нуля) | WholeFromDivision (DIV_ZZ_Z) | Absolute Positivity IntegerFromDivision Addition1 | Морозов
Z-10 | Остаток от деления целого на целое(делитель отличен от нуля) | RemainderFromDivision (MOD_ZZ_Z) | WholeFromDivision Multiplication Subtraction MultiplicationByNegativeOne | Морозов

Номер | Рациональная числа (дроби) — пара (целое; натуральное), первое имеет смысл числителя, второе - знаменателя |   |   | Человек, ответственный за модуль
-- | -- | -- | -- | --
Q-1 | Сокращение дроби | SimplifyingFractions (RED_Q_Q) | Absolute GreatestCommonDivisor WholeFromDivision | Семёнов
Q-2 | Проверка на целое, если рациональное число является целым, то «да», иначе «нет» | CheckingForWhole (INT_Q_B) | | Лицеванова
Q-3 | Преобразование целого в дробное | WholeToFractional (TRANS_Z_Q) | | Грунская
Q-4 | Преобразование дробного в целое (если знаменатель равен 1) | FractionalToWhole (TRANS_Q_Z) | | Грунская
Q-5 | Сложение дробей | Addition (ADD_QQ_Q) | LeastCommonMultiple Multiplication Addition | Комаровский
Q-6 | Вычитание дробей | Subtraction (SUB_QQ_Q) | LeastCommonMultiple Multiplication Subtraction | Комаровский
Q-7 | Умножение дробей | Multiplication (MUL_QQ_Q) | Multiplication | Морозов
Q-8 | Деление дробей (делитель отличен от нуля) | Division (DIV_QQ_Q) | Multiplication | Морозов

Номер | Многочлен с рациональными коэффициентамиm – степень многочлена и массив C коэффициентов |   |   | Человек, ответственный за модуль
-- | -- | -- | -- | --
P-1 | Сложение многочленов | AdditionP(ADD_PP_P) | Addition | Голубев
P-2 | Вычитание многочленов | SubstractionP(SUB_PP_P) | Subtraction | Голубев
P-3 | Умножение многочлена на рациональное число | MultiplicationRational(MUL_PQ_P) | Multiplication | Семёнов
P-4 | Умножение многочлена на x^k | MultipilcationXpowerK(MUL_Pxk_P) | | Голубев
P-5 | Старший коэффициент многочлена | OlderCoeffPoly(LED_P_Q) | | Голубев
P-6 | Степень многочлена | OlderPoly(DEG_P_N) | | Голубев
P-7 | Вынесение из многочлена НОК знаменателей коэффициентов и НОД числителей | GreatestCommonDivisorAndLeastCommonMultipleOfPolynomial(FAC_P_Q) | Absolute FromWholeToNaturals LeastCommonMultiple GreatestCommonDivisor FromNaturalsToWhole WholeFromDivision | Морозов
P-8 | Умножение многочленов | MultiplicationPol(MUL_PP_P) | MUL_PQ_PMUL_Pxk_PADD_PP_P | Грунская
P-9 | Частное от деления многочлена на многочлен при делении с остатком | QuotientOfDivision(DIV_PP_P) | Division DEG_P_NMUL_Pxk_PSUB_PP_PADD_PP_P | Комаровский
P-10 | Остаток от деления многочлена на многочлен при делении с остатком | RemainderFromDivision(MOD_PP_P) | DIV_PP_PMUL_PP_PSUB_PP_P | Комаровский
P-11 | НОД многочленов | GCF_PP_P | GreatestCommonDivisor(DEG_P_NMOD_PP_P) | Пименов
P-12 | Производная многочлена | Derivative(DER_P_P) | | Тростин
P-13 | Преобразование многочлена — кратные корни в простые | SimplifyRoots(NMR_P_P) | GCF_PP_PDER_P_PDIV_PP_P | Тростин

Ответственные за графический интерфейс
-- |
Пименов Голубев Тростин
