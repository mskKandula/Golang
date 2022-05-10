cd to video folder

run the below command with actual filename


ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8


//for various resolutions
ffmpeg -i docker.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 360_out.m3u8

ffmpeg -i docker.mp4 -profile:v baseline -level 3.0 -s 800x480 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 480_out.m3u8

https://dev.to/nodir_dev/transcode-video-source-to-hls-playlist-format-for-video-on-demand-vod-streaming-3h99