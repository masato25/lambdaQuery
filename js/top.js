t2 = _.map(input, function(res){
  res.Avg = _.reduce(res.Values, function(sum,v){
    return (sum+v.Value)
  },0) / (res.Values.length === 0 ? 1 : res.Values.length)
  return res;
}) 

t3 = _.chain(t2).sortBy(function(res){
  return - res.Avg;
}).first(3).value();

console.log("get top 3 reocrds")
_.each(t3, function(x){
  console.log(x.Endpoint, x.Avg)
})
output = JSON.stringify(t3)
