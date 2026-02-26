### Auth Service
Отвечает за аутентификацию и проверку токенов доступа. Сервис предоставляет два эндпоинта:
1. POST /v1/auth/login - принимает учетные данные пользователя и возвращает токен доступа
   #### Ввод верных данных: возвращает код 200 ОК и json с токеном
   <img width="1357" height="524" alt="2026-02-26_15-33-09" src="https://github.com/user-attachments/assets/2a337b76-024a-42d6-bc1a-ac179e824c64" />
   <img width="585" height="139" alt="image" src="https://github.com/user-attachments/assets/a0544ed8-6b6d-4011-aec0-694a73bcdffe" />

   #### Ввод неврных данных: возвращает ошибку 401 Unautorized и json c текстом ошибки
   <img width="1384" height="498" alt="2026-02-26_15-33-31" src="https://github.com/user-attachments/assets/3a611a0c-5a8a-4639-92ef-6591193c9e0d" />
   <img width="592" height="146" alt="image" src="https://github.com/user-attachments/assets/0d5cc2d7-84e1-4566-9170-b922d402a51f" />

3. GET /v1/auth/verify - проверяет валидность токена, и возвращает информацию о пользователе
   #### Проверка верного токена: возвращает код 200 ОК и json с данными о пользователе
   <img width="1377" height="495" alt="2026-02-26_15-33-51" src="https://github.com/user-attachments/assets/0948f716-1bde-4e33-9940-d13703a064db" />
   <img width="607" height="254" alt="image" src="https://github.com/user-attachments/assets/78512256-9c4f-4a2c-b186-28195a4d6308" />

   #### Проверка неверного токенв: возвращает ошибку 401 Unautorized и json c текстом ошибки
   <img width="1372" height="510" alt="2026-02-26_15-34-00" src="https://github.com/user-attachments/assets/17c774b9-8154-4506-a596-289a9f06e584" />
   <img width="577" height="254" alt="image" src="https://github.com/user-attachments/assets/6468bb7e-befc-4c5b-be19-da57961e90f2" />

### Tasks Service
Управляет задачами пользователя. Сервис реализует CRUD операции:
1. POST /v1/tasks - создание новой задачи
   #### Создание новой задачи с авторизацией: возвращает код 201 Created и json с даннымио задаче
   <img width="1369" height="490" alt="2026-02-26_15-34-35" src="https://github.com/user-attachments/assets/367ccc78-3abd-4be8-9700-d22207286766" />
   <img width="1331" height="329" alt="image" src="https://github.com/user-attachments/assets/8ae9a01c-d9f5-419c-a33a-c75ceb913118" />

   #### Попытка сосздать задачу без авторизации: возвращает ошибку 401 Unautorized и json c текстом ошибки
   <img width="1372" height="456" alt="2026-02-26_15-36-12" src="https://github.com/user-attachments/assets/38b46e81-9824-48d3-ab5c-a292e6c5ab78" />
   <img width="789" height="276" alt="image" src="https://github.com/user-attachments/assets/d21c6657-c7d1-4fbe-a842-428990bf2363" />

3. GET /v1/tasks - получение списка всех задач
   #### Получение списка задач: возвращает код 200 ОК и json с информацией обо всех задачах
   <img width="1380" height="615" alt="2026-02-26_15-35-14" src="https://github.com/user-attachments/assets/4ea6f363-13ae-4b87-a51a-6081e372dd8e" />
   <img width="1333" height="170" alt="image" src="https://github.com/user-attachments/assets/ea3b8975-3518-4510-82a6-0e1e867b3640" />

5. GET /v1/tasks/{id} - получение задачи по ID
   #### Получение задачи по существующему ID: возвращает код 200 ОК и json с информацией о задаче
   <img width="1368" height="509" alt="2026-02-26_15-35-29" src="https://github.com/user-attachments/assets/edd5eb2f-2b99-49f6-8716-8e4aa32c0dc9" />
   <img width="1333" height="215" alt="image" src="https://github.com/user-attachments/assets/dfeb95fc-1225-4be0-9d8a-25176c70b19b" />

   #### Попытка получить задачу по несуществующему ID: возвращает ошибку 404 Not Found и json c текстом ошибки
   <img width="1383" height="488" alt="2026-02-26_15-37-54" src="https://github.com/user-attachments/assets/4dddf54c-2e5f-4349-8565-1171dfce62b8" />
   <img width="577" height="190" alt="image" src="https://github.com/user-attachments/assets/3c951825-779b-4274-ac2f-8ea528f2b9aa" />

7. PATCH /v1/tasks/{id} - обновление задачи
   #### Полное обновление задачи: возвращает код 200 ОК и json с информацией о задаче
   <img width="1372" height="482" alt="2026-02-26_15-36-43" src="https://github.com/user-attachments/assets/54ebd10f-2c73-434d-822d-f954f5a638c0" />
   <img width="1341" height="394" alt="image" src="https://github.com/user-attachments/assets/e1e4b2e7-28d6-40b3-bc4c-b8e671466419" />

   #### Частичное обновление задачи: возвращает код 200 ОК и json с информацией о задаче
   <img width="1368" height="485" alt="2026-02-26_15-37-02" src="https://github.com/user-attachments/assets/e590e166-1b25-4afb-aa9e-0ebaae579780" />
   <img width="1327" height="328" alt="image" src="https://github.com/user-attachments/assets/23694812-815b-4706-9517-cd4ee376444d" />
   
9. DELETE /v1/tasks/{id} - удаление задачи
   #### Удаление задачи: возвращает код 204 No Content и json вида
   <img width="1372" height="469" alt="2026-02-26_15-37-40" src="https://github.com/user-attachments/assets/4c15af41-4931-446c-b0ec-1af05f1fe90e" />
   <img width="601" height="169" alt="image" src="https://github.com/user-attachments/assets/a1089127-ffc7-42b4-86b0-aceb4860df18" />

    #### Попытка удалить несуществующую задачу: возвращает ошибку 404 Not Found и json c текстом ошибки
    <img width="1373" height="441" alt="2026-02-26_15-37-46" src="https://github.com/user-attachments/assets/a8bdb68d-3bf8-410f-8a3c-8412971bc59f" />
    <img width="647" height="250" alt="image" src="https://github.com/user-attachments/assets/ed403429-e83f-4e19-b177-66a34fb124a4" />



### Auth Service разворачивается на порту 8081
<img width="393" height="135" alt="image" src="https://github.com/user-attachments/assets/0e6ce13c-eb04-4913-85e7-3445181871c8" />

### Tasks Service разворачивается на порту 8082 и имеет authBaseURL = "http://localhost:8081" для связи с сервисом авторизации
<img width="484" height="248" alt="image" src="https://github.com/user-attachments/assets/3a4b27e4-5623-489b-947a-009e1d1700a4" />







