package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToSlice(t *testing.T) {
	arg := "line1\nline2\nlin'e'3"
	expected := []string{"line1", "line2", "lin'e'3"}
	result := StringToSlice(arg)

	// 使用 reflect.DeepEqual 对比两个切片是否一致
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, result)
	}

	arg = "bash\n/opt/server/prepare_resource.sh\n--config_pb_txt=\"face_detector_type: YOLOV5 face_detector_model_config { ep: CPU onnx_model_file: '/opt/server/model/YoloV5Face.onnx' intra_op_num_threads: 4 inter_op_num_threads: 2 input_count: 1 output_count: 3 } chroma_key_filter_config { similarity: 0.4 smoothness: 0.08 diagonal_coord_coeff: 0.707 chroma_key_red: 0 chroma_key_green: 1 chroma_key_blue: 0 chroma_key_auto: false } preview_avatar_config { frame: 0 top: 170 width: 360 height: 360 } video_config { batch_size: 16 bitrate: 6000 width: 1080 height: 1920 start_frame: 0 end_frame: -1 frames_per_second: 30 }\"\n--face_detector_onnx=\"oss://dev-fantasy/wav2lip/face_detector/YoloV5Face.onnx\"\n--input_video=\"oss://dev-fantasy/wav2lip/Fantasy_avatar/20230316/TT/TT_demoD_speak.mp4\"\n--output=\"oss://dev-fantasy/wav2lip/video_resource_repo/TT_demoD_speak_xxx\""

	expected = []string{
		"bash",
		"/opt/server/prepare_resource.sh",
		"--config_pb_txt=\"face_detector_type: YOLOV5 face_detector_model_config { ep: CPU onnx_model_file: '/opt/server/model/YoloV5Face.onnx' intra_op_num_threads: 4 inter_op_num_threads: 2 input_count: 1 output_count: 3 } chroma_key_filter_config { similarity: 0.4 smoothness: 0.08 diagonal_coord_coeff: 0.707 chroma_key_red: 0 chroma_key_green: 1 chroma_key_blue: 0 chroma_key_auto: false } preview_avatar_config { frame: 0 top: 170 width: 360 height: 360 } video_config { batch_size: 16 bitrate: 6000 width: 1080 height: 1920 start_frame: 0 end_frame: -1 frames_per_second: 30 }\"",
		"--face_detector_onnx=\"oss://dev-fantasy/wav2lip/face_detector/YoloV5Face.onnx\"",
		"--input_video=\"oss://dev-fantasy/wav2lip/Fantasy_avatar/20230316/TT/TT_demoD_speak.mp4\"",
		"--output=\"oss://dev-fantasy/wav2lip/video_resource_repo/TT_demoD_speak_xxx\""}
	result = StringToSlice(arg)
	fmt.Println("len(====): ", result[0])
	fmt.Println("len(====): ", result[1])
	fmt.Println("len(====): ", result[2])
	fmt.Println("len(====): ", result[3])
	fmt.Println("len(====): ", result[4])
	fmt.Println("len(====): ", result[5])
	// 使用 reflect.DeepEqual 对比两个切片是否一致
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, result)
	}
}
