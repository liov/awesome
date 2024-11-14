import cv from '@techstark/opencv-js'
import {Jimp} from 'jimp';
cv.onRuntimeInitialized = async () => {
    const jimpSrc = await Jimp.read(String.raw`xxx.jpg`);
    var src = cv.matFromImageData(jimpSrc.bitmap)

    const dst = new cv.Mat();
    cv.cvtColor(src, dst, cv.COLOR_RGBA2GRAY);
    await new Jimp({
        width: src.cols,
        height: src.rows,
        data: Buffer.from(dst.data)
    }).write('output.jpg');
    src.delete();
    dst.delete();
};
