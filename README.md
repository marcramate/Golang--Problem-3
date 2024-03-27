
# Problem 3: Transaction Broadcasting and Monitoring Client
### Description:
* English Version :
  * develop a module/class interacts with an HTTP Server. This client module will enable the broadcasting of a transaction and subsequently monitor its status until finalization.
  * client module with the following capabilities:
     * Transaction: Construct a JSON payload and send a POST request
     * Status Monitoring 
* Thai Version :
  *  พัฒนา module/class โดยใช้ HTTP Server, โดยจะทำการส่งออกไปพร้อมรับค่ากลับมาเพื่อดู สเตตัส
  * client module มีความสามาถคือ
      *    Transaction ทำการสร้าง payload ในรูปแบบของ JSON และ ส่งค่าไป
      *  ตรวจสอบสถานะ
## Use Language
 * Golang 
## EX.
* Englist Version : 
   * POST
      *    Create variable const URL.
      *  Create struct payload and define type. 
      * Initialize the payload struct and set its fields and convert data to JSON before sending.
      * Create struct Response and define type. 
   * GET
      *    Response from Post add to variable. Then, use variable GET to check status.
      * 5 second delay to Log display status to prevent endless loop.
 * Thai Version :
   * POST
       * สร้างตัวแปร const เพื่อเก็บค่า URL ที่จะส่งไป 
       * สร้างตัวแปรแบบกลุ่ม(struct) ของ payload และ กำหนดtype ของตัวแปรเพื่อนำไปใช้กับการส่งข้อมูล
       * นำ struct ของ payload  ที่สร้างไว้มากำหนดให้ payload พร้อมกับใส่ค่าตาม field ที่กำหนดไว้และทำการ แปลงข้อมูล		  ไปในรูปแบบ JSON ก่อนที่จะส่งค่าไป
       * สร้างตัวแปรแบบกลุ่ม(struct) ของ Response และ กำหนดtype ของตัวแปรเพื่อนำไปใช้ในการรับค่าข้อมูลที่ได้กลับมา
   * GET
       * นำค่าที่ได้จากการ POST กลับมานำใส่ตัวแปรไว้ แล้วนำไป GET ค่าต่อเพื่อนำมาเช็ค status  
       * โดยจะทำการ ดีเลย์ไว้ 5 วินาที ในการโชว์ status ในแต่ละครั้ง เพื่อกันการ loop ไม่สิ้นสุด
  * Status Monitoring  
```
 Status   |  Return  
--------- | ----------
CONFIRMED | processed and confirmed
FAILED 	  | failed to process
PENDING   | awaiting processing
DNE       | does not exist

```


## Result

```
Transaction  xxx.......
Status : PENDING
Transaction is awaiting processing
 ```
 ```
Status : CONFIRMED
Transaction has been processed and confirmed
 ```


