# بازی گل یا پوچ ساده

در بازی گل یا پوچ سه نفر با هم، هم‌تیمی می‌شوند. هر کدام دو دست دارند. دست چپ را با `L` و دست راست را با `R` نشان می‌دهیم. پس در مجموع شش دست در بازی هست. در یک دست گل و پنج دست دیگر پوچ است.

بازیکن‌ها را با اعداد `1` تا `3` شماره‌گذاری می‌کنیم. در ابتدا گل در دست `x` بازیکن شماره‌ی `s` است که مقدار `x` برابر `L` یا `R` است. بازی شروع می‌شود و در حین بازی، بازیکنان `n` حرکت انجام می‌دهند. حرکت‌ها یکی از حالت‌های زیر است:

### خالی بازی k  
در این حرکت بازیکن شماره‌ی `k` محتوای دو دستش را جابه‌جا می‌کند.  
`k ∈ {1, 2, 3}`

### جابه‌جایی k و k+1
در این حرکت بازیکن شماره‌ی `k` محتوای دست `x` خودش را با محتوای دست `y` بازیکن `k+1` جابه‌جا می‌کند.  
`k ∈ {1, 2}` , `x, y ∈ {L, R}`

## ورودی
در سطر اول ورودی،‌ عدد صحیح و مثبت `n` داده می‌شود.  
`(1 ≤ n ≤ 100)`  


در سطر دوم ورودی، عدد `s` و کاراکتر `x` داده می‌شود.  
`(s ∈ {1, 2, 3}, x ∈ {L, R})`  

در `n` سطر بعدی در هر سطر یکی از حرکت‌ها داده می‌شود:
- حرکت نوع اول
- حرکت نوع دوم

## خروجی
در یک سطر، عدد `f` و کاراکتر `y` را چاپ کنید که شماره‌ی بازیکن و دستی که گل در آن است را نشان می‌دهد.

## مثال‌ها

### ورودی نمونه ۱
```text
5
2 L
2 2 R L
2 1 L L
1 1
1 2
1 3
```
### خروجی نمونه ۱
```text
1 R
```
### ورودی نمونه ۲
```text
2
3 R
1 1
2 1 R R
```
### خروجی نمونه ۲
```text
3 R
```

منبع: `کوئرا`