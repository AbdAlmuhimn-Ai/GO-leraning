PS C:\Users\0\Documents\GO\exam04> go version
go version go1.27.1 windows/amd64
PS C:\Users\0\Documents\GO\exam04> go mod tidy
PS C:\Users\0\Documents\GO\exam04> go fmt ./...
report\report_test.go
PS C:\Users\0\Documents\GO\exam04> go vet ./...
PS C:\Users\0\Documents\GO\exam04> go test ./...
?       example.com/academy/exam04      [no test files]
ok      example.com/academy/exam04/rental       (cached)
ok      example.com/academy/exam04/report       0.674s
PS C:\Users\0\Documents\GO\exam04> go run ./exam04
stat C:\Users\0\Documents\GO\exam04\exam04: directory not found
PS C:\Users\0\Documents\GO\exam04> 
PS C:\Users\0\Documents\GO\exam04> go run .
Rented CAM
insufficient equipment
item not found
amount must be positive
Returned CAM
insufficient rented units
CAM 3 1
MIC 2 0
Rental records: CAM=1 MIC=0
90
20
30
140
2
PS C:\Users\0\Documents\GO\exam04> 
PS C:\Users\0\Documents\GO\exam04> git status
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   rental/item.go
        modified:   rental/store.go
        modified:   report/report.go
        modified:   report/report_test.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        ../README.md
        ../answers/level5-unit1-tooling.md

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\Users\0\Documents\GO\exam04>
PS C:\Users\0\Documents\GO\exam04> git --no-pager diff
warning: in the working copy of 'exam04/rental/item.go', LF will be replaced by CRLF the next time Git touches it
warning: in the working copy of 'exam04/rental/store.go', LF will be replaced by CRLF the next time Git touches it
warning: in the working copy of 'exam04/report/report.go', LF will be replaced by CRLF the next time Git touches it
warning: in the working copy of 'exam04/report/report_test.go', LF will be replaced by CRLF the next time Git touches it
diff --git a/exam04/rental/item.go b/exam04/rental/item.go
index a93e7c9..02833fb 100644
--- a/exam04/rental/item.go
+++ b/exam04/rental/item.go
@@ -3,11 +3,11 @@ package rental
 import "errors"

 type Item struct {
-    Code      string
-    Name      string
-    DailyRate int
-    Available int
-    Rented    int
+       Code      string
+       Name      string
+       DailyRate int
+       Available int
+       Rented    int
 }

 func (i Item) Value() int {
@@ -36,4 +36,4 @@ func (i *Item) Return(units int) error {
                i.Rented -= units
                return nil
        }
-}
\ No newline at end of file
+}
diff --git a/exam04/rental/store.go b/exam04/rental/store.go
index 38f921a..3ca03d0 100644
--- a/exam04/rental/store.go
+++ b/exam04/rental/store.go
@@ -25,4 +25,4 @@ func ReturnByCode(items []Item, code string, units int) error {
                return errors.New("item not found")
        }
        return items[idx].Return(units)
-}
\ No newline at end of file
+}
diff --git a/exam04/report/report.go b/exam04/report/report.go
index 02954f7..2236df7 100644
--- a/exam04/report/report.go
+++ b/exam04/report/report.go
@@ -20,4 +20,4 @@ func CountAbove(items []Valuable, threshold int) int {
                }
        }
        return count
-}
\ No newline at end of file
+}
diff --git a/exam04/report/report_test.go b/exam04/report/report_test.go
index e2809df..024e65f 100644
--- a/exam04/report/report_test.go
+++ b/exam04/report/report_test.go
@@ -30,4 +30,21 @@ func TestCountAbove(t *testing.T) {
        if CountAbove(items, 25) != 2 || CountAbove(empty, 25) != 0 {
                t.Fatalf("CountAbove failed for threshold 25 or empty slice")
        }
-}
\ No newline at end of file
+
+}
+
+func TestCountAboveBoundary(t *testing.T) {
+       items := []Valuable{
+               rental.Item{DailyRate: 30, Available: 1},
+               rental.Item{DailyRate: 30, Available: 3},
+       }
+       got30 := CountAbove(items, 30)
+       if got30 != 1 {
+               t.Fatalf("CountAbove(items, 30) = %d; want 1", got30)
+       }
+
+       got90 := CountAbove(items, 90)
+       if got90 != 0 {
+               t.Fatalf("CountAbove(items, 90) = %d; want 0", got90)
+       }
+}
PS C:\Users\0\Documents\GO\exam04>

PS C:\Users\0\Documents\GO\exam04> git --no-pager diff --staged
diff --git a/exam04/rental/item.go b/exam04/rental/item.go
index a93e7c9..02833fb 100644
--- a/exam04/rental/item.go
+++ b/exam04/rental/item.go
@@ -3,11 +3,11 @@ package rental
 import "errors"

 type Item struct {
-    Code      string
-    Name      string
-    DailyRate int
-    Available int
-    Rented    int
+       Code      string
+       Name      string
+       DailyRate int
+       Available int
+       Rented    int
 }

 func (i Item) Value() int {
@@ -36,4 +36,4 @@ func (i *Item) Return(units int) error {
                i.Rented -= units
                return nil
        }
-}
\ No newline at end of file
+}
diff --git a/exam04/rental/store.go b/exam04/rental/store.go
index 38f921a..3ca03d0 100644
--- a/exam04/rental/store.go
+++ b/exam04/rental/store.go
@@ -25,4 +25,4 @@ func ReturnByCode(items []Item, code string, units int) error {
                return errors.New("item not found")
        }
        return items[idx].Return(units)
-}
\ No newline at end of file
+}
diff --git a/exam04/report/report.go b/exam04/report/report.go
index 02954f7..2236df7 100644
--- a/exam04/report/report.go
+++ b/exam04/report/report.go
@@ -20,4 +20,4 @@ func CountAbove(items []Valuable, threshold int) int {
                }
        }
        return count
-}
\ No newline at end of file
+}
diff --git a/exam04/report/report_test.go b/exam04/report/report_test.go
index e2809df..024e65f 100644
--- a/exam04/report/report_test.go
+++ b/exam04/report/report_test.go
@@ -30,4 +30,21 @@ func TestCountAbove(t *testing.T) {
        if CountAbove(items, 25) != 2 || CountAbove(empty, 25) != 0 {
                t.Fatalf("CountAbove failed for threshold 25 or empty slice")
        }
-}
\ No newline at end of file
+
+}
+
+func TestCountAboveBoundary(t *testing.T) {
+       items := []Valuable{
+               rental.Item{DailyRate: 30, Available: 1},
+               rental.Item{DailyRate: 30, Available: 3},
+       }
+       got30 := CountAbove(items, 30)
+       if got30 != 1 {
+               t.Fatalf("CountAbove(items, 30) = %d; want 1", got30)
+       }
+
+       got90 := CountAbove(items, 90)
+       if got90 != 0 {
+               t.Fatalf("CountAbove(items, 90) = %d; want 0", got90)
+       }
+}
PS C:\Users\0\Documents\GO\exam04>
## E — أسئلة التشخيص

1. **الفرق بين git diff و git diff --staged لنفس الملف:**
   لأن `git diff --staged` يعرض التعديلات التي تم تجهيزها بـ `git add` ومقارنتها بآخر Commit، بينما `git diff` يعرض التعديلات الجديدة اللاحقة التي كُتبت في مساحة العمل (Working Directory) ولم تُضف بعد إلى منطقة التجهيز (Staging).

2. **استنتاج نجاح go vet وفشل go test:**
   نستنتج أن الكود نظيف وسليم برمجياً وخالٍ من الأخطاء والعيوب الهيكلية الشائعة، لكنه يحتوي على خطأ منطقي (Logic Error) يخالف المطلوب وظيفياً في نتائج الاختبار.

3. **عدم ظهور go.sum ولماذا لا ننشئه يدويًا:**
   لا يظهر `go.sum` لأن المشروع يعتمد فقط على حزم Go القياسية (Standard Library) أو حزم محلية دون أي مكتبات خارجية. ولا ننشئه يدوياً لأن Go يولّد فيه بصمات تشفيرية (Hashes) دقيقة تلقائياً للتحقق من أمان ونزاهة الحزم ومنع العبث بها.

4. **فائدة مراجعة Diff وفتح Pull Request قبل الدمج إلى main:**
   لضمان مراجعة الكود بدقة (Code Review) واكتشاف الأخطاء مبكراً، والتأكد من نجاح جميع الاختبارات التلقائية قبل دمجها، وذلك لحماية الفرع الرئيسي `main` وإبقائه مستقراً وجاهزاً للإنتاج دائماً.
