avg = 0
t2 = _.map(input, function(res){
  res.Avg = _.reduce(res.Values, function(sum,v){
    return (sum+v.Value)
  },0) / (res.Values.length === 0 ? 1 : res.Values.length)
  avg += res.Avg
  return res;
}) 

avg = avg/t2.length
console.log("current avg number: " + avg)
t3 = _.filter(t2, function(x){
  if(x.Avg > avg){
    return x.Endpoint, x.Avg
  }
})
output = JSON.stringify(t3)
