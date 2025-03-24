import cv2
boxes_mark=[[668,109,1068,509],[667,615,1067,1015],[668,615,1068,1015],[668,615,1068,1015],[668,616,1068,1016]]
boxes_mark_sam=[785,213,971,398]
img=cv2.imread(r"D:\work\xxx.jpg")
for i,box in enumerate(boxes_mark):
    cv2.rectangle(img, (box[0], box[1]), (box[2], box[3]), (0, 0, 50*i+50), 2)
cv2.rectangle(img, (boxes_mark_sam[0], boxes_mark_sam[1]), (boxes_mark_sam[2], boxes_mark_sam[3]), (0, 255, 0), 2)
cv2.imwrite("rect.jpg",img)
print((boxes_mark[0][0]+boxes_mark[0][2])//2,(boxes_mark_sam[0]+boxes_mark_sam[2])//2)