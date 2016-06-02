/*
  get avg value of counter and retrun top n
*/
limit = (typeof limit == "undefined"? 3 : limit)
orderby = (typeof orderby == "undefined"? "desc" : orderby)
t2 = _.map(input, function(res){
  res.Avg = _.reduce(res.Values, function(sum,v){
    return (sum+v.Value)
  },0) / (res.Values.length === 0 ? 1 : res.Values.length)
  return res;
})

t3 = _.chain(t2).sortBy(function(res){
  if(orderby == "desc"){
    return - res.Avg
  }else{
    return res.Avg
  }
}).first(limit).value();

output = JSON.stringify(t3)
