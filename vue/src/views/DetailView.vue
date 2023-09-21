<template>
    <div>
        <div style="width: 100%; height: auto;">
            <el-tabs v-model="activeName" @tab-click="handleClick" tab-position="top" style="width: 100%;">
                <el-tab-pane label="用户管理" name="first">用户管理</el-tab-pane>
                <el-tab-pane label="配置管理" name="second">配置管理</el-tab-pane>
                <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>
                <el-tab-pane label="定时任务补偿" name="fourth">定时任务补偿</el-tab-pane>
            </el-tabs>
        </div>

        <el-row>
            <el-col :span="6"><img src="@/assets/孙尚香时之恋人.jpg" alt="时之恋人" width="100px"></el-col>
            <el-col :span="6"><img src="@/assets/孙尚香时之恋人.jpg" alt="时之恋人" width="100px"></el-col>
            <el-col :span="6"><img src="@/assets/孙尚香时之恋人.jpg" alt="时之恋人" width="100px"></el-col>
            <el-col :span="6"><img src="@/assets/孙尚香时之恋人.jpg" alt="时之恋人" width="100px"></el-col>
        </el-row>

        <!-- 按钮 -->
        <el-row>
            <el-col :span="24">
                <el-button>默认按钮</el-button>
                <el-button type="primary">主要按钮</el-button>
            </el-col>
            <el-col :span="24">
                <el-button plain>默认按钮</el-button>
                <el-button type="primary" plain>主要按钮</el-button>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="24">
                <el-button circle icon="el-icon-edit" type="success">编辑按钮</el-button>
                <el-button plain icon="el-icon-delete" type="danger"></el-button>
            </el-col>
        </el-row>

        <!-- 输入框 -->
        <el-row style="margin-top: 20px">
            <el-col>
                <el-input show-password style="width: 200px;" v-model="input" placeholder="请输入内容"
                    type="textarea"></el-input>
                <el-input style="width: 200px;" v-model="input2" placeholder="请输入内容"
                    prefix-icon="el-icon-search"></el-input>
                <el-input style="width: 200px;" v-model="input3" placeholder="请输入内容" suffix-icon="el-icon-user"></el-input>
                <el-input show-password style="width: 200px;" v-model="password" placeholder="密码"></el-input>
            </el-col>
        </el-row>

        <el-row style="margin: 20px 0;">
            <el-col>
                <el-input clearable style="width: 200px;" v-model="input4" placeholder="请输入内容"
                    prefix-icon="el-icon-search"></el-input>
                <el-button>搜索一下</el-button>

            </el-col>
        </el-row>

        <!-- 带搜索建议的搜索框 -->
        <el-row style="margin: 20px 0;">
            <el-col>
                <el-autocomplete :fetch-suggestions="querySearch" :trigger-on-focus="true" placeholder="请输入内容"
                    style="width: 400px;" v-model="inputQuery"></el-autocomplete>
            </el-col>
        </el-row>

        <!-- 下拉框 -->
        <el-row>
            <el-select clearable v-model="fruitSelected" @change="changeSelect">
                <el-option value="苹果"></el-option>
                <el-option value="香蕉"></el-option>
                <el-option value="梨子"></el-option>
            </el-select>

            <el-select clearable v-model="fruitSelected" @change="changeSelect">
                <el-option v-for="item in fruits" :key="item.id" :label="item.value" :value="item.value"></el-option>
            </el-select>

            <el-select clearable multiple v-model="userSelected" @change="changeUser">
                <el-option v-for="item in users" :key="item.card" :label="item.name" :value="item.card"></el-option>
            </el-select>

            <el-select clearable filterable v-model="userSelected2" @change="changeUser">
                <el-option v-for="item in users2" :key="item.card" :label="item.label" :value="item.card"></el-option>
            </el-select>
        </el-row>

        <!-- 选项 -->
        <el-row>
            <el-col :span="24">
                <el-radio v-model="radio1" label="1">备选项1</el-radio>
                <el-radio v-model="radio1" label="2">备选项2</el-radio>
            </el-col>
            <el-col :span="24">
                <el-radio-group v-model="radio2" @change="selectRadio">
                    <el-radio label="1">备选项1</el-radio>
                    <el-radio label="2">备选项2</el-radio>
                    <el-radio label="3">备选项3</el-radio>
                    <el-radio label="4">备选项4</el-radio>
                </el-radio-group>
            </el-col>
            <el-col :span="24">
                <el-radio-group v-model="radio3">
                    <el-radio-button label="1">备选项1</el-radio-button>
                    <el-radio-button label="2">备选项2</el-radio-button>
                    <el-radio-button label="3">备选项3</el-radio-button>
                    <el-radio-button label="4">备选项4</el-radio-button>
                </el-radio-group>
            </el-col>

            <el-checkbox-group v-model="radio4" @change="selectCheckBox">
                <el-checkbox label="复选框1"></el-checkbox>
                <el-checkbox label="复选框2"></el-checkbox>
                <el-checkbox label="复选框3"></el-checkbox>
            </el-checkbox-group>
        </el-row>

        <!-- 日期和时间 -->
        <el-date-picker v-model="date" type="date" placeholder="选择日期" value-format="yyyy-MM-dd"
            @change="changeDate"></el-date-picker>
        <el-date-picker v-model="dateTime" type="datetime" placeholder="选择日期时间" value-format="yyyy-MM-dd hh:mm:ss"
            @change="changeDateTime"></el-date-picker>
        <el-date-picker v-model="dateRange" type="datetimerange" start-placeholder="选择开始时间" end-placeholder="选择结束时间"
            value-format="yyyy-MM-dd hh:mm:ss" @change="changeDateRange" style="width: 400px;"></el-date-picker>

        <!-- 表格 -->
        <el-row style="margin: 20px 0;">
            <el-table :data="tableData" :header-cell-style="{ background: 'lightblue', fontSize: '20px' }" stripe border>
                <el-table-column label="名字" prop="name" align="center"></el-table-column>
                <el-table-column label="编号" prop="id" align="center"></el-table-column>
                <el-table-column label="年龄" prop="age" align="center"></el-table-column>
            </el-table>
        </el-row>

    </div>
</template>



<script>
export default {
    data() {
        return {
            radio1: "1",
            radio2: "1",
            radio3: "1",
            radio4: [],
            input: "",
            input2: "",
            input3: "",
            input4: "",
            password: "",
            inputQuery: "",
            fruits: [
                { value: "1苹果", id: "1" }, { value: "2香蕉", id: "2" }, { value: "3梨子", id: "3" }
            ],
            fruitSelected: "",
            users: [
                { name: "hhh", card: "111" },
                { name: "www", card: "222" },
                { name: "uuu", card: "333" },
            ],
            userSelected: "",
            users2: [
                { label: "hhh", card: "111" },
                { label: "www", card: "222" },
                { label: "uuu", card: "333" },
            ],
            userSelected2: "",
            date: "",
            dateTime: "",
            dateRange: "",
            tableData: [
                { name: "hhh", id: "111", age: "10" },
                { name: "www", id: "222", age: "20" },
                { name: "uuu", id: "333", age: "30" },
                { name: "ttt", id: "444", age: "40" },
            ],
        }
    },
    methods: {
        querySearch(query, cb) {
            let result = query ? this.fruits.filter(v => v.value.includes(query)) : this.fruits;
            cb(result);
        },
        changeSelect() { },
        changeUser() {
            console.log(this.userSelected);
        },
        selectRadio() {
            alert(this.radio2);
        },
        selectCheckBox() {
            console.log(this.radio4);
        },
        changeDate() {
            console.log(this.date)
        },
        changeDateTime() {
            console.log(this.dateTime)
        },
        changeDateRange() {
            console.log(this.dateRange)
        },
    }
}
</script>