cd to video folder

run the below command with actual filename


ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8


//for various resolutions
ffmpeg -i docker.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 360_out.m3u8

ffmpeg -i docker.mp4 -profile:v baseline -level 3.0 -s 800x480 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 480_out.m3u8

https://dev.to/nodir_dev/transcode-video-source-to-hls-playlist-format-for-video-on-demand-vod-streaming-3h99


https://gist.github.com/Andrey2G/78d42b5c87850f8fbadd0b670b0e6924

ffmpeg -i golang.mp4 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -c:v libx264 -crf 22 -c:a aac -ar 48000 -filter:v:0 scale=w=480:h=360 -maxrate:v:0 600k -b:a:0 64k -filter:v:1 scale=w=640:h=480 -maxrate:v:1 900k -b:a:1 128k -filter:v:2 scale=w=1280:h=720 -maxrate:v:2 1500k -b:a:2 128k -var_stream_map "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p" -preset slow -hls_list_size 0 -threads 0 -f hls -hls_playlist_type event -hls_time 10 -hls_flags independent_segments -master_pl_name "master.m3u8" -y "%v/index.m3u8"