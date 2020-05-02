<html>
	<body>
		<h1>Final Project Group18</h1>

		<p>Please login using your email and password</p>
		<form action="<?php echo htmlspecialchars($_SERVER["PHP_SELF"]); ?>" method="get">
			<p><input type="text" name="email" /></p>
			<p><input type="text" name="password" /></p>
			<p><input type="submit" /></p>
		</form>

		<?php
			$p = '/(a+)+$/';
			if(isset($_GET['email']) && isset($_GET['password'])){
				$matched = false;
				$email = $_GET['email'];
				$password = $_GET['password'];
				if((!empty($email)) && (!empty($password))){
				    if(strlen($email) > 25){
				        echo "<p>Your input too long!!!</p>";
				    }else{
				        if(preg_match($p, $email)){
                        	$matched=true;
               			}
                        if ($matched){
                        	echo "<p>Hello {$email} </p>";
                        }else{
                        	echo "<p>Ops, invalid email </p>";
                        }
				    }
				}else{
					echo "<p>Value cannot be empty!!!</p>";
				}
			}
		?>
	</body>
</html>




