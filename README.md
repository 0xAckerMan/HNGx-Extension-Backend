# HNGX Chrome Extension Backend

### Video screen recording extension backend
#### Endpoints available
- **Healthcheck for the api** ```/api/health``` ***[Test](https://chrome-extension-e75d.onrender.com/api/health)***
```
https://chrome-extension-e75d.onrender.com/api/health
```
![](https://hackmd.io/_uploads/H1FdXfdxa.png)


- **Posting or uploading a video**  ```/uploader/:videoname```</br>
Takes videos as chunks and stores them in the disk under the uploads directory

    example:
```
curl -X POST -H "X-Video-Filename: 'VIDEO NAME'"  -T  "PATH OF THE VIDEO" https://chrome-extension-e75d.onrender.com/uploader/NAME TO BE SAVED WITH
```
![](https://hackmd.io/_uploads/SJ5nmfOxT.png)

- **Listing all videos** ```/videos``` ***[test](https://chrome-extension-e75d.onrender.com/videos)***</br> 
Lists all videos that have been uploaded using the ```GET``` method
```curl https://chrome-extension-e75d.onrender.com/videos | jq```

![img](https://hackmd.io/_uploads/BJ_mQzdga.png)

- **Getting a single video** ```url/videos/name``` [test](https://chrome-extension-e75d.onrender.com/videos/submission_vid.mp4)</br>
Used to play or view a single upladed or saved video from the server

- **Deleting a recording** </br>
This takes the video name you want to delete and deletes it 
```curl -X DELETE https://chrome-extension-e75d.onrender.com/videos/videoname```
![](https://hackmd.io/_uploads/S1zmSGOgp.png)


#### Comming soon
- Transcription
