package gomodels

type TestA struct {
	MyFieldA1 string `field:"my_field_a1" gen:"add:true;update:true;delete:true;query:true"`
	MyFieldA2 string `field:"my_field_a2" gen:"add:true;update:false;delete:true;query:true"`
	MyFieldA3 string `field:"my_field_a3" gen:"add:true;update:false;delete:false;query:true"`
}

type TestB struct {
	MyFieldB1 string `field:"my_field_b1" gen:"add:true;update:true;delete:true;query:false"`
	MyFieldB2 string `field:"my_field_b2" gen:"add:true;update:false;delete:true;query:false"`
	MyFieldB3 string `field:"my_field_b3" gen:"add:false;update:false;delete:false;query:true"`
}
