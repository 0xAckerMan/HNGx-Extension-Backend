# HNGX Chrome Extension Backend

### Video screen recording extension backend
#### Endpoints available
- **Healthcheck for the api** ```/api/health``` ***[Test](https://chrome-extension-e75d.onrender.com/api/health)***
```
https://chrome-extension-e75d.onrender.com/api/health
```
![](https://hackmd-prod-images.s3-ap-northeast-1.amazonaws.com/uploads/upload_4408cfc21438ce4fb83a1eb39f3182a1.png?AWSAccessKeyId=AKIA3XSAAW6AWSKNINWO&Expires=1696241663&Signature=DkChmJv7%2F2IHU6i8R6u0BsmnHQs%3D)


- **Posting or uploading a video**  ```/uploader/:videoname```</br>
Takes videos as chunks and stores them in the disk under the uploads directory

    example:
```
curl -X POST -H "X-Video-Filename: 'VIDEO NAME'"  -T  "PATH OF THE VIDEO" https://chrome-extension-e75d.onrender.com/uploader/NAME TO BE SAVED WITH
```
![](https://hackmd-prod-images.s3-ap-northeast-1.amazonaws.com/uploads/upload_f70aebc403d3d96560287a92b1483092.png?AWSAccessKeyId=AKIA3XSAAW6AWSKNINWO&Expires=1696241754&Signature=8tZyQqVkRH%2BStW6kqL2Q2CLqkAs%3D)

- **Listing all videos** ```/videos``` ***[test](https://chrome-extension-e75d.onrender.com/videos)***</br> 
Lists all videos that have been uploaded using the ```GET``` method
```curl https://chrome-extension-e75d.onrender.com/videos | jq```

![img](https://hackmd-prod-images.s3-ap-northeast-1.amazonaws.com/uploads/upload_2f9193da42faf9d997961a09a9d84166.png?AWSAccessKeyId=AKIA3XSAAW6AWSKNINWO&Expires=1696241800&Signature=E5rKi65hgORoYwjdxkixPXufXiY%3D)

- **Getting a single video** ```url/videos/name``` [test](https://chrome-extension-e75d.onrender.com/videos/submission_vid.mp4)</br>
Used to play or view a single upladed or saved video from the server

- **Deleting a recording** </br>
This takes the video name you want to delete and deletes it 
```curl -X DELETE https://chrome-extension-e75d.onrender.com/videos/videoname```
![](https://hackmd-prod-images.s3-ap-northeast-1.amazonaws.com/uploads/upload_d3c6cf38d82e7a75596256d350168511.png?AWSAccessKeyId=AKIA3XSAAW6AWSKNINWO&Expires=1696241838&Signature=YZh6Qys2GKcRCJQrxc1j5dyaGbs%3D)


#### Comming soon
- Transcription
