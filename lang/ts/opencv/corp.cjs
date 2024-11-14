const {Jimp} = require('jimp');

async function onRuntimeInitialized() {
    //console.log(cv.getBuildInformation())
    // load local image file with jimp. It supports jpg, png, bmp, tiff and gif:
//使用jimp加载本地镜像文件。它支持jpg，png，bmp，tiff和gif：
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
}


Module = {
    onRuntimeInitialized
}

const cv = require('./opencv.cjs');

