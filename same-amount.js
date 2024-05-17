
function sameAmount(str, reg1, reg2) {
     return execReg(str, reg1) === execReg(str, reg2)
}

function execReg(str, reg) {
     let count = 0;
     let match;
     while ((match = reg.exec(str)) !== null) {
          count++;
          str = str.slice(match.index + 1)
     }
     return count;
}
