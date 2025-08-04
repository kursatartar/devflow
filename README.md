# DevFlow v0.2

## Proje Tanımı
Bu projede Go diliyle user, project ve task gibi kavramları nasıl oluşturup yöneteceğimi denedim.  
Tüm veriler bellekte tutuluyor. Dosya ya da veritabanı kullanımı yok.  
Amaç sadece öğrendiğim Go kavramlarını uygulamak ve geliştirmekti.

## Özellikler
- User, Project ve Task oluşturma, listeleme, güncelleme ve silme işlemleri
- Task'ların bir projeye bağlı olacak şekilde tanımlanması (ProjectID alanı üzerinden)
- Struct yapıları ve map kullanımı
- Varlık kontrolü (`value, ok := map[key]`)
- `Profile` gibi iç içe struct kullanımı
- Status bilgileri için `const` sabitleri (`StatusPending`, `StatusActive`, vs.)
- JSON tag’leri yorum satırı olarak eklendi (ileride API’ye dönüşebilir)

## Proje Yapısı

cmd/

└── main.go // Uygulamanın giriş noktası

models/

├── user.go // User struct'ı ve map

├── project.go // Project struct'ı ve map

└── task.go // Task struct'ı ve map

handlers/

├── user_handler.go // User CRUD işlemleri

├── project_handler.go // Project CRUD işlemleri

└── task_handler.go // Task CRUD işlemleri


## Versiyonlar

### v0.1
- `map[string]string` ile sadece `id -> name` tutularak user işlemleri denendi
- CRUD işlemleri sadece user için yapıldı
- Her şey basit fonksiyonlarla yönetildi

### v0.2
- Struct kullanımı eklendi (User, Project, Task)
- Map’ler `map[string]User`, `map[string]Project`, vs. olarak güncellendi
- Project ve Task yapıları detaylandırıldı
- Task’lara ProjectID ilişkisi eklendi
- `main.go` üzerinden elle test yapılarak geliştirildi

## Dikkat Edilen Noktalar
- Update işlemlerinde struct map’ten alınıp güncellenip tekrar map’e atandı
- Map erişimlerinde `if _, ok := map[id]; ok` şeklinde kontrol kullanıldı
- Tüm ID’ler elle veriliyor, otomatik ID üretimi yapılmadı
- Aynı ID ile tekrar create edilirse eski veri üzerine yazılır
- `projectID`, `assignedTo` gibi alanlar string olarak tutuldu çünkü `"p1"`, `"u2"` gibi okunabilir ID’ler kullanıldı
- Bellekte tutulduğu için uygulama kapanınca tüm veriler silinir
- Projede şu an sadece Task → Project arasında ilişki var

## Ek Notlar
- Input’lar kullanıcıdan alınmıyor, `main.go` üzerinden manuel fonksiyon çağrılarıyla test ediliyor
- JSON tag’leri şimdilik yorum satırına eklendi, ileride API'ye geçilirse aktif edilebilir
- Kod sade, fonksiyonel, öğrenim odaklı tutuldu  
