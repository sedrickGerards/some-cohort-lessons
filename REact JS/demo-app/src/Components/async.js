// const sum = new Promise((resolve, reject) => {
//   let number = 1 + 2

//   if (number === 3){
//     resolve("success")
//   } else{
//     reject("Failed")
//   }

// })

// sum.then((message) => {
//   console.log(message)
// }).catch((message) => {
//   console.log(message);
  
// })

const userWatchingVideo = new Promise((resolve, reject) => {

  let userLeft = false;
  let userWacthingMeme = false;

  if (userLeft){
    reject('user left')
  } else if (userWacthingMeme){
    reject({
      name: "User is watching memes 😂",
      });
    
  } else{
    resolve('hureey user is watching')
    
  }

})

userWatchingVideo.then((message) =>{
  console.log("resolved:", message)
}).catch((err) =>{
  console.log("rejected:",err);
  
})