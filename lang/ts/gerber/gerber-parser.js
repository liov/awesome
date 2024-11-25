import fs from 'fs'
import gerberParser from 'gerber-parser'

var parser = gerberParser({})

parser.on('warning', function(w) {
  console.warn('warning at line ' + w.line + ': ' + w.message)
})

fs.createReadStream( String.raw`xxx`)
  .pipe(parser)
  .on('data', function(obj) {
    console.log(JSON.stringify(obj))
  })