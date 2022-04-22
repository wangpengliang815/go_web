<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="UserId" width="95">
        <template slot-scope="scope">
          {{  scope.row.Id }}
        </template>
      </el-table-column>
      <el-table-column label="UserName" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.UserName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Age" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.Age }}
        </template>
      </el-table-column>
      <el-table-column label="PhoneNumber" align="center">
        <template slot-scope="scope">
          {{ scope.row.PhoneNumber }}
        </template>
      </el-table-column>
      <el-table-column label="Address"  align="center">
        <template slot-scope="scope">
          {{ scope.row.Address }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getList } from '@/api/user'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.list = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
