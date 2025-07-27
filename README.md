# devflowv2

bu projede go ile çalışırken user, project ve task gibi şeyleri nasıl oluşturup yöneteceğimi denedim.  
bütün verileri bellekte tuttum. database yok, dosya yazma yok.  
her şey map ve struct'lar üzerinden dönüyor.

projenin amacı sadece öğremdiklerimi uygulamak.  
başta map kullandım, sonra struct öğrendim, onları entegre ettim.  
şimdilik sadece main.go'dan manuel olarak fonksiyonları çağırıp test ediyorum.

## proje yapısı

cmd/  
→ main.go burada. uygulama buradan başlatılıyor.

models/  
→ user, project ve task struct'ları burada tanımlı  
→ ayrıca onları tuttuğumuz map'ler de burada

handlers/  
→ user_handler.go: user işlemleri  
→ project_handler.go: project işlemleri  
→ task_handler.go: task işlemleri  
→ hepsi create, list, update, delete içeriyor

## ilk versiyon (v0.1 diyebilirim)

başta map[string]string kullanarak user'ları tuttum  
örneğin id -> name gibi  
crud işlemlerini handlers/user_handler.go içinde yazdım  
main.go dosyasından çağırarak test ettim

## v0.2'de yaptıklarım

go'da struct yapısını öğrendim  
user'ı artık struct olarak tanımladım: id, name, email

map[string]User yapısına geçtik  
crud fonksiyonları struct'a göre güncellendi  
aynı şekilde project ve task için de struct'lar oluşturdum

task struct'ı içinde ProjectID alanı var  
bu sayede bir görevin hangi projeye ait olduğunu biliyoruz  
task'ı oluştururken geçerli bir proje id verilmesi gerekiyor

## dikkat ettiğim bazı şeyler

- update işlemlerinde struct'ı map'ten alıp alanlarını değiştirdim, sonra geri map'e koydum
- map erişiminde `value, ok := map[key]` şeklinde kontrol yaptım. (value yerine "_" de kullanabilirim çünkü burada bana lazım olan value değil, sadece ok değerinin true veya false dönmesi çünkü böyle bir değer var mı diye kontrol etmek istiyorum.)
- ikinci değer (ok) daima bool olur. true = key var, false = yok. go dilinde bir kural.





## diğer notlar

- repo'da her değişiklikten sonra terminale git status yazdım, burada hangi dosyalarda değişiklik yaptığımı görerek sonrasında "git add" komutu ile pushlayacağım dosyaları ekledim. sonrasında git commit -m "commit notu" ve push kullanarak kodumu repomda güncelledim.
- bellekte veri tuttuğum için uygulama kapanınca her şey silinir. kalıcılık yok.
- user, project ve task’lar arasında şu an yalnızca task → project ilişkisi var
- tüm id’ler elle veriliyor. otomatik id üretimi yapılmadı
- aynı id ile tekrar create yaparsam eskisinin üstüne yazar, uyarı verilmez  
- input kullanıcıdan alınmıyor, main.go'dan manuel test yapılıyor
- project ve task id'lerini string olarak tuttum çünkü "p1", "p2" gibi isimler kullandım  eğer sadece sayı kullansaydım (1, 2, 3...), o zaman int daha mantıklı olurdu
ama string olması şu an hem esnek hem okunabilir oldu, işimi kolaylaştırdı  
- Project adında bir struct tanımladım, her projenin id, name ve description bilgisi var
- Projects adında bir map oluşturdum, string id'leri key olarak kullanıyorum
ve map'in içindeki her değer bir Project struct'ı yani tüm proje bilgilerini içeriyor. bu sayede birden fazla projeyi tek yerde tutabiliyorum ve id ile erişiyorum

