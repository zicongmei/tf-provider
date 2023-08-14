output "last_updated" {
  value = item1.default.last_updated
}

output "uuid" {
  value = item1.default.foo[0].bar[0].uuid
}

output "id" {
  value = item1.default.id
}


output "created" {
  value = item1.default.created
}