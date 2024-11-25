
import fs from 'fs'
import pcbStackup from 'pcb-stackup'

const fileNames = [
  String.raw`xxx`,
]

const layers = fileNames.map(filename => ({
  filename,
  gerber: fs.createReadStream(filename),
}))

pcbStackup(layers).then(stackup => {
  console.log(stackup.top.svg) // logs "<svg ... </svg>"
  console.log(stackup.bottom.svg) // logs "<svg ... </svg>"
})