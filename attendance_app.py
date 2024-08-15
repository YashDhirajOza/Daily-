import streamlit as st
import cv2
import numpy as np
import tempfile
import os

@st.cache_resource
def load_yolo():
    current_dir = os.path.dirname(os.path.abspath(__file__)) 
    weights_path = os.path.join(current_dir, "yolov3.weights")
    config_path = os.path.join(current_dir, "yolov3.cfg")
    
    net = cv2.dnn.readNet(weights_path, config_path)
    net.setPreferableBackend(cv2.dnn.DNN_BACKEND_CUDA)
    net.setPreferableTarget(cv2.dnn.DNN_TARGET_CUDA)
    layer_names = net.getLayerNames()
    output_layers = [layer_names[i - 1] for i in net.getUnconnectedOutLayers()]
    return net, output_layers

def detect_people(frame, net, output_layers):
    height, width, _ = frame.shape
    blob = cv2.dnn.blobFromImage(frame, 1/255.0, (416, 416), swapRB=True, crop=False)
    
    net.setInput(blob)
    outs = net.forward(output_layers)
    
    people_boxes = []
    confidences = []
    for out in outs:
        for detection in out:
            scores = detection[5:]
            class_id = np.argmax(scores)
            confidence = scores[class_id]
            if confidence > 0.5 and class_id == 0:  
                center_x, center_y, w, h = (detection[0:4] * np.array([width, height, width, height])).astype('int')
                x = int(center_x - w / 2)
                y = int(center_y - h / 2)
                people_boxes.append([x, y, int(w), int(h)])
                confidences.append(float(confidence))
    
    indices = cv2.dnn.NMSBoxes(people_boxes, confidences, 0.5, 0.4)
    return [people_boxes[i] for i in indices.flatten()]

def draw_outline(frame, people_boxes):
    mask = np.zeros_like(frame)
    for box in people_boxes:
        x, y, w, h = box
        mask[y:y+h, x:x+w] = frame[y:y+h, x:x+w]
    
    gray = cv2.cvtColor(mask, cv2.COLOR_BGR2GRAY)
    _, thresh = cv2.threshold(gray, 1, 255, cv2.THRESH_BINARY)
    contours, _ = cv2.findContours(thresh, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    
    cv2.drawContours(frame, contours, -1, (0, 255, 0), 2)
    return frame

# Streamlit App
st.title('Automatic Student Counter in Video')

uploaded_file = st.file_uploader("Choose a video file", type=['mp4', 'avi'])

if uploaded_file is not None:
    net, output_layers = load_yolo()
    with tempfile.NamedTemporaryFile(delete=False, suffix='.mp4') as temp_file:
        temp_file.write(uploaded_file.read())
        temp_file_path = temp_file.name
    cap = cv2.VideoCapture(temp_file_path)
    fps = int(cap.get(cv2.CAP_PROP_FPS))
    width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    
    frame_count = 0
    total_students = 0
    video_placeholder = st.empty()
    while cap.isOpened():
        ret, frame = cap.read()
        if not ret:
            break
        
        frame_count += 1
        
        people_boxes = detect_people(frame, net, output_layers)
        current_students = len(people_boxes)
        total_students += current_students
        
        frame = draw_outline(frame, people_boxes)
        
        cv2.putText(frame, f'Students: {current_students}', (10, 30), 
                    cv2.FONT_HERSHEY_SIMPLEX, 1, (0, 255, 0), 2)
        frame_rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
        video_placeholder.image(frame_rgb, channels="RGB", use_column_width=True)
        cv2.waitKey(50)
        
    cap.release()
    os.remove(temp_file_path)
    avg_students = total_students // frame_count
    st.write(f"Estimated average number of students per frame: {avg_students}")
