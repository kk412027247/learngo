const str = 'abcabcbb';
const str1 = 'bbbbb';
const str2 = 'pwwkew';

const find = str=>{
  const arr = Array(str.length).fill('').map((item,index)=>str[index]);
  const shortArr = [...new Set(arr)];
  console.log(shortArr)
  const shortArrLength = shortArr.length;
  console.log(shortArrLength);
  arr.reduce((pre,cur,_arr,index)=>{
    if(index > _arr.length - shortArrLength){
      return pre
    }else{
      
    }
  },[])




  // return arr.reduce((pre,cur,index,arr)=>{
  //   //   const subArr = Array(index+1).fill('').map((item,index)=>arr[index]);
  //   //   const shortArr = [... new Set(subArr)];
  //   //   const length =  shortArr.length;
  //   //   if(pre.length < length){
  //   //     return{
  //   //       length:length,
  //   //       sub: shortArr
  //   //     }
  //   //   }else{
  //   //     return pre
  //   //   }
  //   // },{length:0,sub:[]})
};
find(str2)

// console.log(find(str).sub.toString());
// console.log(find(str1).sub.toString());
// console.log(find(str2).sub.toString());
