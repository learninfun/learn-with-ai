

1. 如何在 Terraform 中聲明自定義的 VPC？
2. 如何使用 Terraform 動態地創建 EC2 實例？
3. 如何在 Terraform 中定義 Auto Scaling？
4. 如何在 Terraform 中實現密碼加密？
5. 如何在 Terraform 中定義 Lambda 函數並綁定 CloudWatch 觸發器？

答案：

1. 
resource "aws_vpc" "custom_vpc" {
        cidr_block = "10.0.0.0/16"
        instance_tenancy = "dedicated"
}

2.
resource "aws_instance" "example_ec2" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  dynamic "ebs_block_device" {
    for_each = var.enabled_block_device ? [1] : []
    content {
      device_name = "/dev/sdh"
      volume_size = 20
      delete_on_termination = true
    }
  }

  tags = {
    Name = "ExampleInstance"
  }
}

3.
resource "aws_autoscaling_group" "example_asg" {
  name                 = "example"
  desired_capacity     = 2
  max_size             = 2
  min_size             = 2
  default_cooldown     = 300
  health_check_grace_period = 3600
  health_check_type    = "EC2"

  launch_configuration = aws_launch_configuration.example_lc.id

  vpc_zone_identifier  = [aws_subnet.example_subnet.id]
  target_group_arns = [aws_alb_target_group.example_tg.arn]

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

4.
data "aws_ssm_parameter" "example" {
  name = "example_parameter"
  with_decryption = true
}

5.
resource "aws_lambda_function" "example_lambda" {
  filename      = "lambda_function_payload.zip"
  function_name = "example_lambda"
  role          = aws_iam_role.example_role.arn

  handler = "lambda_function_payload.handler"
  runtime = "nodejs12.x"
  timeout = 60

  environment {
    variables = {
      EXAMPLE_VAR = "example"
    }
  }
}

resource "aws_cloudwatch_event_rule" "example_rule" {
  name        = "example_rule"
  description = "Example rule"
  schedule_expression = "rate(1 minute)"
}

resource "aws_cloudwatch_event_target" "example_target" {
  target_id = "example_target"
  rule      = aws_cloudwatch_event_rule.example_rule.name
  input     = jsonencode({
    lambda_name = aws_lambda_function.example_lambda.function_name
  })
  arn = aws_lambda_function.example_lambda.arn
}
