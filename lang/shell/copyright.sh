/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

 # windows的find会覆盖,打开msys2自带终端执行
find /path/to/your/directory -type f -name "*.go" -exec sed -i '1i/*\n * Copyright 2024 hopeio. All rights reserved.\n * Licensed under the MIT License that can be found in the LICENSE file.\n * @Created by jyb\n */\n' {} \;