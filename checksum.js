/**
 * Calculating the Checksum for FIX PROTOCOL messages using NodeJS
 *
 * Example messages
 * 8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|
 * Reference : https://stackoverflow.com/questions/32708068/how-to-calculate-checksum-in-fix-manually
 */

var message = process.argv[2];

checksum = 0;
for (let index in message) {
  if (message[index] == '|') {
    checksum += 1;
  } else {
    checksum += message[index].charCodeAt(0);
  }
}

checksum = checksum % 256;
checksum = ('00' + checksum).slice(-3);
message = message + '10=' + checksum + '|';
console.table({
  Input: process.argv[2],
  Output: message,
});
