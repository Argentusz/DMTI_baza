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
P-8 | Умножение многочленов | MultiplicationPol(MUL_PP_P) | MultiplicationRational MultipilcationXpowerK AdditionP | Грунская
P-9 | Частное от деления многочлена на многочлен при делении с остатком | QuotientOfDivision(DIV_PP_P) | Division OlderPoly MultipilcationXpowerK SubstractionP AdditionP | Комаровский
P-10 | Остаток от деления многочлена на многочлен при делении с остатком | RemainderFromDivision(MOD_PP_P) | QuotientOfDivision MultiplicationPol SubstractionP | Комаровский
P-11 | НОД многочленов | GreatestCommonDivisor(GCF_PP_P) | OlderPoly RemainderFromDivision | Пименов
P-12 | Производная многочлена | Derivative(DER_P_P) | | Тростин
P-13 | Преобразование многочлена — кратные корни в простые | SimplifyRoots(NMR_P_P) | GreatestCommonDivisor Derivative QuotientOfDivision | Тростин

Ответственные за графический интерфейс
-- |
Голубев Тростин

# Расходы времени в модуле полиномов

![cpuAdditionP-1](https://user-images.githubusercontent.com/75137969/163597883-79a0d215-8535-4845-80ea-c01dfdbc4054.png)
![cpuCompare-1](https://user-images.githubusercontent.com/75137969/163621337-cd078c55-d6e8-4c56-b3c3-618e04343f2f.png)
![cpuCopyP-1](https://user-images.githubusercontent.com/75137969/163621367-200732ca-097d-43a5-9648-d0bdbc6ac274.png)
![cpuDerivative-1](https://user-images.githubusercontent.com/75137969/163621425-8422c517-ebb9-45fa-a04d-aa8f0d9bf9d4.png)
![cpuGreatestCommonDivisor-1](https://user-images.githubusercontent.com/75137969/163621488-85f4378d-a20b-41af-a59a-717e789b1079.png)
![cpuGreatestCommonDivisorAndLeastCommonMultipleOfPolynomial-1](https://user-images.githubusercontent.com/75137969/163621553-b2cad0bf-20eb-4f77-b285-433cd9cd8684.png)
![cpuMultiplicationPol-1](https://user-images.githubusercontent.com/75137969/163621571-e99aaf44-30d4-4030-9e72-878086d821cb.png)
![cpuMultiplicationRational-1](https://user-images.githubusercontent.com/75137969/163621579-da0554f1-b289-4764-aa60-44178aa453f6.png)
![cpuMultiplicationXpowerK-1](https://user-images.githubusercontent.com/75137969/163621589-3787821a-bcb1-46b5-834a-46d94bdf6549.png)
![cpuOlderCoeffPoly-1](https://user-images.githubusercontent.com/75137969/163621598-a84576bd-0247-4e34-a135-fceba0677ce9.png)
![cpuOlderPoly-1](https://user-images.githubusercontent.com/75137969/163621650-b1a6b87f-55cf-4ff2-a177-345237ef7804.png)
![cpuPolynomial_MakeP-1](https://user-images.githubusercontent.com/75137969/163621668-e8e9ad79-6972-4c81-9052-189db0128087.png)
![cpuPolynomial_ToStringPol-1](https://user-images.githubusercontent.com/75137969/163621686-758111e4-c39c-47c0-9769-aa9ef5f69923.png)
![cpuQuotientOfDivision-1](https://user-images.githubusercontent.com/75137969/163621698-82410dd6-ece4-40be-a797-46f8cfbf6554.png)
![cpuRemainderFromDivision-1](https://user-images.githubusercontent.com/75137969/163621710-dc553119-8176-43d6-ae48-d7bdac7d463c.png)
![cpuSubstractionP-1](https://user-images.githubusercontent.com/75137969/163621722-344c216b-c479-442b-abea-7bd54d5ecefa.png)

# Расходы памяти в модуле полиномов

![memAdditionP-1](https://user-images.githubusercontent.com/75137969/163624074-11493ba4-da04-4f20-a5c7-26b92f5e0a04.png)
![memCompare-1](https://user-images.githubusercontent.com/75137969/163624080-a29ef41f-49f1-4d22-83ef-8f82923b819d.png)
![memCopyP-1](https://user-images.githubusercontent.com/75137969/163624095-dacae7b8-909e-49bd-8afd-8591f6d5cc97.png)
![memDerivative-1](https://user-images.githubusercontent.com/75137969/163624110-998ceb76-9c48-46b9-ad43-5f1bdf67387d.png)
![memGreatestCommonDivisor-1](https://user-images.githubusercontent.com/75137969/163624139-17e07341-9287-4c92-aa61-0a12654c5936.png)
![memGreatestCommonDivisorAndLeastCommonMultipleOfPolynomial-1](https://user-images.githubusercontent.com/75137969/163624161-2f99c0bf-b6a3-4157-9785-c8898fdb282e.png)
![memMultiplicationPol-1](https://user-images.githubusercontent.com/75137969/163624171-2ff2305e-c284-4e4c-8c52-5bba84232644.png)
![memMultiplicationRational-1](https://user-images.githubusercontent.com/75137969/163624185-33d90f7f-f714-4059-8a28-a477e72ab2fc.png)
![memMultiplicationXpowerK-1](https://user-images.githubusercontent.com/75137969/163624195-52ab62fe-fc3c-4639-a8e8-7ff2ac1489cb.png)
![memOlderCoeffPoly-1](https://user-images.githubusercontent.com/75137969/163624208-058865c2-b58c-4379-b456-ed75cd407cd9.png)
![memOlderPoly-1](https://user-images.githubusercontent.com/75137969/163624217-049cd011-44e2-464c-a424-256b10925d2d.png)
![memPolynomial_MakeP-1](https://user-images.githubusercontent.com/75137969/163624226-7b09b1f1-f217-4103-9971-7e71727ac3a7.png)
![memPolynomial_ToStringPol-1](https://user-images.githubusercontent.com/75137969/163624232-f073a14c-d987-4ba8-a55c-c76954f9bfa4.png)
![memQuotientOfDivision-1](https://user-images.githubusercontent.com/75137969/163624236-7df9d89f-ec39-4516-b732-5e15344a8d48.png)
![memRemainderFromDivision-1](https://user-images.githubusercontent.com/75137969/163624275-07c37660-4d2f-4d08-ac71-2f63d0e6c5a2.png)
![memSubstractionP-1](https://user-images.githubusercontent.com/75137969/163624300-c8322433-84d0-44e7-bc08-262c1a44461e.png)

