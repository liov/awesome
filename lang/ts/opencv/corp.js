import cv from 'opencv.js'
import {Jimp} from 'jimp'
async function onRuntimeInitialized() {

        // load local image file with jimp. It supports jpg, png, bmp, tiff and gif:
//使用jimp加载本地镜像文件。它支持jpg，png，bmp，tiff和gif：
    const jimpSrc = await Jimp.read('./lena.jpg');
    cv.matFromImageData(jimpSrc.bitmap).then(image => {
            const grayImage = new cv.Mat();
            cv.cvtColor(image, grayImage, cv.COLOR_RGBA2GRAY);
        new Jimp({
            width: grayImage.cols,
            height: grayImage.rows,
            data: Buffer.from(grayImage.data)
        }).write('output.png')
        });
}